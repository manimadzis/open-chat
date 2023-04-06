package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type userRepository struct {
	pool *pgx.ConnPool
}

func (u userRepository) Create(ctx context.Context, user entities.User) (entities.UserId, error) {
	sql := `INSERT INTO 
			"user"(login, password, registration_time, nickname)
			VALUES($1, $2, $3, $4)
			RETURNING id`
	row := u.pool.QueryRowEx(ctx, sql, nil,
		user.Login,
		user.Password,
		user.RegistrationTime,
		user.Nickname,
	)
	err := row.Scan(&user.Id)
	if err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pq.ErrorCode(pgErr.Code).Name() == "unique_violation" && pgErr.ConstraintName == "user_login_key" {
				return 0, services.ErrLoginAlreadyExists
			}
		}
		return 0, services.NewUnknownError(err)
	}
	return user.Id, nil
}

func (u userRepository) FindById(ctx context.Context, userId entities.UserId) (*entities.User, error) {
	sql := `SELECT login, password, registration_time, nickname
			FROM "user"
			WHERE id = $1`
	user := entities.User{Id: userId}
	row := u.pool.QueryRowEx(ctx, sql, nil,
		userId,
	)
	err := row.Scan(
		&user.Login,
		&user.Password,
		&user.RegistrationTime,
		&user.Nickname,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, services.ErrNoSuchUser
		}
		return nil, services.NewUnknownError(err)
	}
	return &user, nil
}

func (u userRepository) FindByLogin(ctx context.Context, login string) (*entities.User, error) {
	sql := `SELECT id, password, registration_time, nickname
			FROM "user"
			WHERE login = $1`
	user := entities.User{Login: login}
	row := u.pool.QueryRowEx(ctx, sql, nil,
		login,
	)
	err := row.Scan(
		&user.Id,
		&user.Password,
		&user.RegistrationTime,
		&user.Nickname,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, services.ErrNoSuchLogin
		}
		return nil, services.NewUnknownError(err)
	}

	return &user, nil
}

func NewUserRepository(pool *pgx.ConnPool) services.UserRepository {
	return &userRepository{pool: pool}
}
