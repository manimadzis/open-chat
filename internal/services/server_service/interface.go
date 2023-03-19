package server_service

import (
	"context"
	"open-chat/internal/entities"
)

type ServerService interface {
	Create(ctx context.Context, server *entities.Server, user *entities.User) error
	Delete(ctx context.Context, server *entities.Server, user *entities.User) error
	Join(ctx context.Context, server *entities.Server, user *entities.User) error
	Kick(ctx context.Context, server *entities.Server, user *entities.User) error
}
