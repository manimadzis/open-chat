package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type channelRepo struct {
	pool *pgx.ConnPool
}

func (c channelRepo) Create(ctx context.Context, channel entities.Channel) (entities.ChannelId, error) {
	// TODO implement me
	panic("implement me")
}

func (c channelRepo) Delete(ctx context.Context, channelId entities.ChannelId) error {
	// TODO implement me
	panic("implement me")
}

func (c channelRepo) FindByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Channel, error) {
	// TODO implement me
	panic("implement me")
}

func NewChannelRepository(pool *pgx.ConnPool) services.ChannelRepository {
	return &channelRepo{pool: pool}
}
