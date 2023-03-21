package role_checker

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories"
	"open-chat/internal/services/role_system"
)

type roleChecker struct {
	userRepo repositories.UserRepository
}

func (r roleChecker) Check(ctx context.Context, userId entities.UserId, serverId entities.ServerId, permissions ...entities.PermissionValue) error {
	serverProfile, err := r.userRepo.FindServerProfileByIds(ctx, userId, serverId)
	if err != nil {
		return err
	}
	return role_system.NewRoleSystem(serverProfile.Roles).Check(permissions...)
}

func NewRoleChecker(repository repositories.UserRepository) RoleChecker {
	return &roleChecker{
		userRepo: repository,
	}
}
