package permission_checker

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type permissionChecker struct {
	serverProfileRepo services.ServerProfileRepository
	roleSystem        services.RoleSystem
}

func (r permissionChecker) Check(
	ctx context.Context,
	userId entities.UserId,
	serverId entities.ServerId,
	permissions ...entities.PermissionValue,
) error {
	serverProfile, err := r.serverProfileRepo.FindById(
		ctx,
		entities.ServerProfileId{
			UserId:   userId,
			ServerId: serverId,
		},
	)
	if err != nil {
		return err
	}
	r.roleSystem.SetRoles(serverProfile.Roles)

	return r.roleSystem.Check(permissions...)
}

func NewPermissionChecker(
	serverProfileRepo services.ServerProfileRepository,
	roleSystem services.RoleSystem,
) services.PermissionChecker {
	return &permissionChecker{
		serverProfileRepo: serverProfileRepo,
		roleSystem:        roleSystem,
	}
}
