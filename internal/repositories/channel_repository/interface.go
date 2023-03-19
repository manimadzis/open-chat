package channel_repository

import (
	"context"
	"open-chat/internal/entities"
)

type ChannelRepository interface {
	Create(ctx context.Context, channel *entities.Channel, user *entities.User) error
	Delete(ctx context.Context, channel *entities.Channel, user *entities.User) error
	FindByServer(ctx context.Context, server *entities.Server, user *entities.User) ([]*entities.Channel, error)
}