package server_profile_checker

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type serverProfileChecker struct {
	serverRepo services.ServerRepository
	roleSystem services.RoleSystem
}

func (r serverProfileChecker) Check(
	ctx context.Context,
	userId entities.UserId,
	serverId entities.ServerId,
	permissions ...entities.PermissionValue,
) error {
	serverProfile, err := r.serverRepo.FindServerProfileByIds(
		ctx,
		serverId,
		userId,
	)
	if err != nil {
		return err
	}
	r.roleSystem.SetRoles(serverProfile.Roles)

	return r.roleSystem.Check(permissions...)
}

func NewServerProfileChecker(
	repository services.ServerRepository,
	roleSystem services.RoleSystem,
) services.ServerProfileChecker {
	return &serverProfileChecker{
		serverRepo: repository,
		roleSystem: roleSystem,
	}
}
