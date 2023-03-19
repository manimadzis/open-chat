package message_repository

import (
	"context"
	"open-chat/internal/entities"
)

type MessageRepository interface {
	Create(ctx context.Context, message *entities.Message) error
	Delete(ctx context.Context, message *entities.Message) error
	FindByChannel(ctx context.Context, channel *entities.Channel, offset, count uint64) ([]*entities.Message, error)
}
