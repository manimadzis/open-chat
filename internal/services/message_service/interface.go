package message_service

import (
	"context"
	"open-chat/internal/entities"
)

type MessageService interface {
	Send(ctx context.Context, message *entities.Message) error
	Delete(ctx context.Context, messageId entities.MessageId, userId entities.UserId) error
	FindInChat(ctx context.Context, userId entities.UserId, channelId entities.ChannelId, filters *entities.MessageFiltersDTO) ([]entities.Message, error)
}
