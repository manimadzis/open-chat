package message_service

import (
	"context"
	"open-chat/internal/entities"
)

type MessageService interface {
	Send(ctx context.Context, message *entities.Message) error
	Delete(ctx context.Context, message *entities.Message) error
	Find(ctx context.Context, user *entities.User, channel *entities.Channel, filters *entities.MessageFiltersDTO) ([]*entities.Message, error)
}
