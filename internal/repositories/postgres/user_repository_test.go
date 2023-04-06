package postgres_test

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/postgres"
	"testing"
	"time"
)

func connect() *pgx.ConnPool {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "openchat",
			User:     "postgres",
			Password: "postgres",
		},
	},
	)
	if err != nil {
		panic(err)
	}
	return pool
}

func clear(pool *pgx.ConnPool) {
	sql := `TRUNCATE "user" CASCADE`
	if _, err := pool.Exec(sql); err != nil {
		panic(err)
	}
}

func clearAll(pool *pgx.ConnPool) {
	sql := `DROP TABLE "user" CASCADE`
	pool.Exec(sql)
}

func create(pool *pgx.ConnPool) {
	sql := `CREATE TABLE IF NOT EXISTS "user"
(
    id                BIGSERIAL PRIMARY KEY,
    login             TEXT      NOT NULL UNIQUE,
    password          TEXT      NOT NULL,
    registration_time TIMESTAMP NOT NULL,
    nickname          TEXT      NOT NULL
)`
	if _, err := pool.Exec(sql); err != nil {
		panic(err)
	}
}

func TestUserRepository_Create(t *testing.T) {
	pool := connect()
	// clearAll(pool)
	// create(pool)
	repo := postgres.NewUserRepository(pool)
	id, err := repo.Create(context.Background(), entities.User{
		Login:            "1223",
		Password:         "123",
		Nickname:         "321",
		RegistrationTime: time.Now(),
	},
	)

	require.Equal(t, nil, err)
	require.Equal(t, entities.UserId(1), id)
}

func TestUserRepository_FindById(t *testing.T) {
	pool := connect()
	// clearAll(pool)
	// create(pool)
	repo := postgres.NewUserRepository(pool)
	user, err := repo.FindById(context.Background(), entities.UserId(1))

	require.Equal(t, nil, err)
	fmt.Printf("%#v", user)
	// require.Equal(t, entities.UserId(1), id)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	pool := connect()
	// clearAll(pool)
	// create(pool)
	repo := postgres.NewUserRepository(pool)
	user, err := repo.FindByLogin(context.Background(), "1223")

	require.Equal(t, nil, err)
	fmt.Printf("%#v", user)
	// require.Equal(t, entities.UserId(1), id)
}
