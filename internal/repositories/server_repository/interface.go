package server_repository

import (
	"context"
	"open-chat/internal/entities"
)

type ServerRepository interface {
	Create(ctx context.Context, server *entities.Server, user *entities.User) error
	Delete(ctx context.Context, server *entities.Server, user *entities.User) error
	Join(ctx context.Context, server *entities.Server, user *entities.User) error
}
