package user_service

import (
	"context"
	"open-chat/internal/entities"
)

type UserService interface {
	Create(ctx context.Context, user *entities.User) error
}
