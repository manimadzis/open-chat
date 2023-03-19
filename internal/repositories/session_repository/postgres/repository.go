package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"open-chat/internal/entities"
	"open-chat/internal/repositories"
	"open-chat/internal/repositories/session_repository"
)

type sessionRepository struct {
	pool *pgx.ConnPool
}

func (s *sessionRepository) Create(ctx context.Context, session *entities.Session) error {
	sql := "INSERT INTO session(token, expired_at, user_id) VALUES ($1, $2, $3) RETURNING id"
	err := s.pool.QueryRowEx(
		ctx,
		sql,
		nil,
		session.Token, session.ExpiredAt, session.User.Id).Scan(&session.Id)
	if err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pq.ErrorCode(pgErr.Code).Name() == "unique_violation" && pgErr.ConstraintName == "session_user_id_key" {
				return session_repository.ErrUserAlreadyHasToken
			} else if pq.ErrorCode(pgErr.Code).Name() == "unique_violation" && pgErr.ConstraintName == "session_token_key" {
				return session_repository.ErrTokenAlreadyExists
			}
		}
		return repositories.UnknownError(err)
	}

	return nil
}

func (s *sessionRepository) DeleteByToken(ctx context.Context, session *entities.Session) error {
	sql := "DELETE FROM session WHERE token = $1"
	ct, err := s.pool.ExecEx(ctx, sql, nil, session.Token)
	if err != nil {
		return repositories.UnknownError(err)
	}
	if ct.RowsAffected() == 0 {
		return session_repository.ErrNoSuchToken
	}
	return nil
}

func (s *sessionRepository) FindByToken(ctx context.Context, session *entities.Session) error {
	sql := "SELECT id, user_id, expired_at FROM session WHERE token = $1"
	row := s.pool.QueryRowEx(ctx, sql, nil, session.Token)

	if err := row.Scan(&session.Id, &session.User.Id, &session.ExpiredAt); err != nil {
		if err == pgx.ErrNoRows {
			return session_repository.ErrNoSuchToken
		}
		return repositories.UnknownError(err)
	}
	return nil
}

func (s *sessionRepository) FindByUser(ctx context.Context, session *entities.Session) error {
	panic("not implemented")
}

func NewSessionRepository(pool *pgx.ConnPool) session_repository.SessionRepository {
	return &sessionRepository{
		pool: pool,
	}
}
