package postgres_test

import (
	"github.com/jackc/pgx"
	"open-chat/internal/repositories/postgres"
	"testing"
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

func TestSessionRepository_Create(t *testing.T) {
	pool := connect()
	repo := postgres.NewSessionRepository(pool)

}
