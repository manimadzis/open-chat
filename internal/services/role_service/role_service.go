package role_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"open-chat/internal/services/role_system"
)

type roleService struct {
	roleRepo          services.RoleRepository
	permissionChecker services.PermissionChecker
}

func (r *roleService) Create(
	ctx context.Context,
	role entities.Role,
) (entities.RoleId, error) {
	if err := r.permissionChecker.Check(
		ctx,
		role.CreatorId,
		role.ServerId,
		role_system.PERM_CREATE_ROLE,
	); err != nil {
		return 0, err
	}
	return r.roleRepo.Create(ctx, role)
}

func (r *roleService) Delete(
	ctx context.Context,
	roleId entities.RoleId,
	userId entities.UserId,
	serverId entities.ServerId,
) error {
	if err := r.permissionChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_DELETE_ROLE,
	); err != nil {
		return err
	}
	return r.roleRepo.Delete(ctx, roleId)
}

func (r *roleService) Change(
	ctx context.Context,
	role entities.Role,
	userId entities.UserId,
	serverId entities.ServerId,
) error {
	if err := r.permissionChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_CHANGE_ROLE,
	); err != nil {
		return err
	}
	return r.roleRepo.Change(ctx, role)
}

func (r *roleService) FindByServer(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) ([]entities.Role, error) {
	if err := r.permissionChecker.Check(
		ctx,
		userId,
		serverId,
	); err != nil {
		return nil, err
	}
	return r.roleRepo.FindRolesByServerId(ctx, serverId)
}

func NewRoleService(
	repository services.RoleRepository,
	permissionChecker services.PermissionChecker,
) services.RoleService {
	return &roleService{
		roleRepo:          repository,
		permissionChecker: permissionChecker,
	}
}
