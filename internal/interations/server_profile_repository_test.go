package interations

import (
	"context"
	"log"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/postgres"
	"testing"
)

func TestServerProfileRepository_FindById(t *testing.T) {
	pool := connect()
	repo := postgres.NewServerProfileRepository(pool)
	sp, err := repo.FindById(context.Background(), entities.ServerProfileId{
		UserId:   111,
		ServerId: 1,
	},
	)
	log.Println(err)
	log.Println(sp)
}
