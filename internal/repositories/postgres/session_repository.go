package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type sessionRepository struct {
	pool *pgx.ConnPool
}

func (s *sessionRepository) Create(ctx context.Context, session entities.Session) error {
	sql := "INSERT INTO session(token, expired_at, user_id) VALUES ($1, $2, $3) RETURNING id"
	_, err := s.pool.ExecEx(
		ctx,
		sql,
		nil,
		session.Token, session.ExpiredAt, session.UserId)
	if err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pq.ErrorCode(pgErr.Code).Name() == "unique_violation" && pgErr.ConstraintName == "session_user_id_key" {
				return services.ErrUserAlreadyHasToken
			} else if pq.ErrorCode(pgErr.Code).Name() == "unique_violation" && pgErr.ConstraintName == "session_token_key" {
				return services.ErrTokenAlreadyExists
			}
		}
		return services.UnknownError(err)
	}

	return nil
}

func (s *sessionRepository) DeleteByToken(ctx context.Context, token entities.SessionToken) error {
	sql := "DELETE FROM session WHERE token = $1"
	ct, err := s.pool.ExecEx(ctx, sql, nil, token)
	if err != nil {
		return services.UnknownError(err)
	}
	if ct.RowsAffected() == 0 {
		return services.ErrNoSuchToken
	}
	return nil
}

func (s *sessionRepository) FindByToken(ctx context.Context, token entities.SessionToken) (*entities.Session, error) {
	sql := "SELECT user_id, expired_at FROM session WHERE token = $1"
	row := s.pool.QueryRowEx(ctx, sql, nil, token)
	session := entities.Session{Token: token}
	if err := row.Scan(&session.UserId, &session.ExpiredAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, services.ErrNoSuchToken
		}
		return nil, services.UnknownError(err)
	}
	return &session, nil
}

func (s *sessionRepository) FindByUserId(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	panic("not implemented")
}

func NewSessionRepository(pool *pgx.ConnPool) services.SessionRepository {
	return &sessionRepository{
		pool: pool,
	}
}
