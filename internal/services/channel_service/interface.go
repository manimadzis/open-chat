package channel_service

import (
	"context"
	"open-chat/internal/entities"
)

type ChannelService interface {
	Create(ctx context.Context, channel *entities.Channel, userId entities.UserId) error
	Delete(ctx context.Context, channelId entities.ChannelId, userId entities.UserId) error
	FindByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Channel, error)
}
