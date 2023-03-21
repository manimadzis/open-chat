package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/role_repository/postgres"
	"open-chat/internal/services/role_system"
	"testing"
	"time"
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

func TestRoleRepository_Create(t *testing.T) {
	pool, err := Connect()

	if err != nil {
		t.Error(err)
	}

	repo := postgres.NewRoleRepository(pool)
	role := &entities.Role{
		Name:       "Name",
		Permission: role_system.PERM_SEND_MESSAGE,
		Server:     &entities.Server{Id: 123},
		CreatedBy:  &entities.User{Id: 123},
		CreatedAt:  time.Now(),
	}

	if err = repo.Create(context.Background(), role); err != nil {
		t.Error(err)
	}
}

func TestRoleRepository_Delete(t *testing.T) {
	pool, err := Connect()

	if err != nil {
		t.Error(err)
	}

	repo := postgres.NewRoleRepository(pool)
	role := &entities.Role{
		Id: 1,
	}

	if err = repo.Delete(context.Background(), role); err != nil {
		t.Error(err)
	}
}

func TestRoleRepository_FindByServer(t *testing.T) {
	pool, err := Connect()

	if err != nil {
		t.Error(err)
	}

	repo := postgres.NewRoleRepository(pool)
	role := &entities.Role{
		Server: &entities.Server{Id: 123},
	}

	roles, err := repo.FindRoleByServer(context.Background(), role.Server)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", roles)
}
