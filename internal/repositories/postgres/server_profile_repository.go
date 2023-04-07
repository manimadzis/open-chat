package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type serverProfileRepo struct {
	pool *pgx.ConnPool
}

func (s serverProfileRepo) Create(
	ctx context.Context,
	serverProfile entities.ServerProfile,
) (entities.ServerProfileId, error) {
	serverProfileId := entities.ServerProfileId{
		UserId:   serverProfile.UserId,
		ServerId: serverProfile.ServerId,
	}
	sql := `INSERT INTO role(
				server_id,
				user_id,
				join_time,
				nickname
			VALUES($1, $2, $3, $4)`
	_, err := s.pool.ExecEx(ctx, sql, nil,
		serverProfile.ServerId,
		serverProfile.UserId,
		serverProfile.JoinTime,
		serverProfile.Nickname,
	)

	if err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			codeName := pq.ErrorCode(pgErr.Code).Name()
			if codeName == "foreign_key_violation" &&
				pgErr.ConstraintName == "server_profile_server_id_fkey" {
				return serverProfileId, services.ErrNoSuchServer
			} else if codeName == "foreign_key_violation" &&
				pgErr.ConstraintName == "server_profile_user_id_fkey" {
				return serverProfileId, services.ErrNoSuchUser
			} else if codeName == "unique_violation" &&
				pgErr.ConstraintName == "server_profile_pkey" {
				return serverProfileId, services.ErrServerProfileAlreadyExists
			}
		}
		return serverProfileId, services.NewUnknownError(err)
	}

	return serverProfileId, nil
}

func (s serverProfileRepo) Delete(
	ctx context.Context,
	serverProfileId entities.ServerProfileId,
) error {
	sql := "DELETE FROM server_profile WHERE server_id = $1 AND user_id = $2"
	ct, err := s.pool.ExecEx(ctx, sql, nil,
		serverProfileId.ServerId,
		serverProfileId.UserId,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}
	if ct.RowsAffected() != 1 {
		return services.ErrNoSuchServerProfile
	}
	return nil
}

func (s serverProfileRepo) FindById(
	ctx context.Context,
	serverProfileId entities.ServerProfileId,
) (*entities.ServerProfile, error) {
	sql := `SELECT 
				server_id,
				user_id,
				join_time,
				nickname
			FROM server_profile
			WHERE server_id = $1 AND user_id = $2`

	row := s.pool.QueryRowEx(ctx, sql, nil,
		serverProfileId.ServerId,
		serverProfileId.UserId,
	)

	var serverProfile entities.ServerProfile
	if err := row.Scan(
		&serverProfile.ServerId,
		&serverProfile.UserId,
		&serverProfile.JoinTime,
		&serverProfile.Nickname,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, services.ErrNoSuchServerProfile
		}

		return nil, services.NewUnknownError(err)
	}

	return &serverProfile, nil
}

func NewServerProfileRepository(pool *pgx.ConnPool) services.ServerProfileRepository {
	return &serverProfileRepo{pool: pool}
}
