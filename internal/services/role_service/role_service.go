package role_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/role_repository"
	"open-chat/internal/services/role_checker"
	"open-chat/internal/services/role_system"
)

type roleService struct {
	roleRepo    role_repository.RoleRepository
	roleChecker role_checker.RoleChecker
}

func (r *roleService) Create(ctx context.Context, role *entities.Role, user *entities.User) error {
	if err := r.roleChecker.Check(ctx, user, role.Server, role_system.PERM_CREATE_ROLE); err != nil {
		return err
	}
	return r.roleRepo.Create(ctx, role, user)
}

func (r *roleService) Delete(ctx context.Context, role *entities.Role, user *entities.User) error {
	if err := r.roleChecker.Check(ctx, user, role.Server, role_system.PERM_DELETE_ROLE); err != nil {
		return err
	}
	return r.roleRepo.Delete(ctx, role, user)
}

func (r *roleService) Change(ctx context.Context, role *entities.Role, user *entities.User) error {
	if err := r.roleChecker.Check(ctx, user, role.Server, role_system.PERM_CHANGE_ROLE); err != nil {
		return err
	}
	return r.roleRepo.Change(ctx, role, user)
}

func (r *roleService) FindByServer(ctx context.Context, server *entities.Server, user *entities.User) ([]*entities.Role, error) {
	if err := r.roleChecker.Check(ctx, user, server, role_system.PERM_CHANGE_ROLE); err != nil {
		return []*entities.Role{}, err
	}
	return r.roleRepo.FindByServer(ctx, server, user)
}

func NewRoleService(repository role_repository.RoleRepository, roleChecker role_checker.RoleChecker) RoleService {
	return &roleService{
		roleRepo:    repository,
		roleChecker: roleChecker,
	}
}
