package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"testing"
)

func Connect() (*pgx.ConnPool, error) {
	return pgx.NewConnPool(pgx.ConnPoolConfig{
		MaxConnections: 20,
		ConnConfig: pgx.ConnConfig{
			User:     "postgres",
			Password: "postgres",
			Port:     5432,
			Database: "openchat",
		},
	},
	)
}

func TestSessionRepository_Create(t *testing.T) {
	pool, err := Connect()

	if err != nil {
		t.Error(err)
	}
	_, err = pool.Exec(`CREATE TABLE IF NOT EXISTS session(
		id BIGSERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL,
		token TEXT NOT NULL,
		expired_at TIMESTAMP
		)
	`)
	if err != nil {
		t.Error(err)
	}

	repo := NewSessionRepository(pool)
	session := entities.NewSession(&entities.User{
		Id: 12334,
	})
	session.Token = "5b0e4d81-c56f-4256-832c-7a5dc1dac33a"
	if err = repo.Create(context.Background(), session); err != nil {
		t.Error(err)
	}
}

func TestSessionRepository_Delete(t *testing.T) {
	pool, err := Connect()
	if err != nil {
		t.Error(err)
	}

	session := &entities.Session{
		Token: "ef4c27c5-d6bf-43ee-bda1-f117fec1f2ca",
	}

	repo := NewSessionRepository(pool)
	if err := repo.DeleteByToken(context.Background(), session); err != nil {
		t.Error(err)
	}
}

func TestSessionRepository_FindUser(t *testing.T) {
	pool, err := Connect()
	if err != nil {
		t.Error(err)
	}
	repo := NewSessionRepository(pool)

	session := entities.NewSession(&entities.User{})
	session.Token = "b9fee129-1702-4167-bdcf-6f2f2778f6f0"

	if err := repo.FindByToken(context.Background(), session); err != nil {
		t.Error(err)
	}
}
