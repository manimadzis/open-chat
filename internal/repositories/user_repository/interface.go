package user_repository

import (
	"context"
	"open-chat/internal/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindById(ctx context.Context, user *entities.User) error
	FindServerProfileByIds(ctx context.Context, user *entities.User, server *entities.Server) (*entities.ServerProfile, error)
}
