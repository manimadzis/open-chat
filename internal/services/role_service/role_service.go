package role_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories"
	"open-chat/internal/services/role_checker"
	"open-chat/internal/services/role_system"
)

type roleService struct {
	roleRepo    repositories.RoleRepository
	roleChecker role_checker.RoleChecker
}

func (r *roleService) Create(ctx context.Context, role *entities.Role, userId entities.UserId, serverId entities.ServerId) error {
	if err := r.roleChecker.Check(ctx, userId, serverId, role_system.PERM_CREATE_ROLE); err != nil {
		return err
	}
	return r.roleRepo.Create(ctx, role)
}

func (r *roleService) Delete(ctx context.Context, roleId entities.RoleId, userId entities.UserId, serverId entities.ServerId) error {
	if err := r.roleChecker.Check(ctx, userId, serverId, role_system.PERM_DELETE_ROLE); err != nil {
		return err
	}
	return r.roleRepo.Delete(ctx, roleId)
}

func (r *roleService) Change(ctx context.Context, role *entities.Role, userId entities.UserId, serverId entities.ServerId) error {
	if err := r.roleChecker.Check(ctx, userId, serverId, role_system.PERM_CHANGE_ROLE); err != nil {
		return err
	}
	return r.roleRepo.Change(ctx, role)
}

func (r *roleService) FindByServer(ctx context.Context, serverId entities.ServerId, userId entities.UserId) ([]entities.Role, error) {
	if err := r.roleChecker.Check(ctx, userId, serverId, role_system.PERM_CHANGE_ROLE); err != nil {
		return []entities.Role{}, err
	}
	return r.roleRepo.FindRolesByServerId(ctx, serverId)
}

func NewRoleService(repository repositories.RoleRepository, roleChecker role_checker.RoleChecker) RoleService {
	return &roleService{
		roleRepo:    repository,
		roleChecker: roleChecker,
	}
}
