package user_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/user_repository"
)

type userService struct {
	userRepo user_repository.UserRepository
}

func (u userService) Create(ctx context.Context, user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
