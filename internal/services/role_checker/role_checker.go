package role_checker

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/user_repository"
	"open-chat/internal/services/role_system"
)

type roleChecker struct {
	userRepo user_repository.UserRepository
}

func (r roleChecker) Check(ctx context.Context, user *entities.User, server *entities.Server, permissions ...entities.Permission) error {
	serverProfile, err := r.userRepo.FindServerProfileByIds(ctx, user, server)
	if err != nil {
		return err
	}
	return role_system.NewRoleSystem(serverProfile.Roles).Check(permissions...)
}

func NewRoleChecker(repository user_repository.UserRepository) RoleChecker {
	return &roleChecker{
		userRepo: repository,
	}
}
