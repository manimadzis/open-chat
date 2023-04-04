package role_system

import (
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type roleSystem struct {
	totalPermission entities.PermissionValue
}

func (r *roleSystem) Check(permissions ...entities.PermissionValue) error {
	sumPermission := entities.PermissionValue(0)
	for _, permission := range permissions {
		sumPermission |= permission
	}

	if r.totalPermission&sumPermission != sumPermission {
		return services.ErrNotEnoughPermissions
	}
	return nil
}

func (r *roleSystem) SetRoles(roles []entities.Role) {
	r.totalPermission = entities.PermissionValue(0)

	for _, role := range roles {
		r.totalPermission |= role.PermissionValue
	}
}

func NewRoleSystem(roles []entities.Role) services.RoleSystem {
	rs := &roleSystem{}
	rs.SetRoles(roles)
	return rs
}
