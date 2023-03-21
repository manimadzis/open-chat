package role_system

import "open-chat/internal/entities"

type roleSystem struct {
	totalPermission entities.PermissionValue
}

func (r roleSystem) Check(permissions ...entities.PermissionValue) error {
	sumPermission := entities.PermissionValue(0)
	for _, permission := range permissions {
		sumPermission |= permission
	}

	if r.totalPermission&sumPermission != sumPermission {
		return ErrNotEnoughPermissions
	}
	return nil
}

func NewRoleSystem(roles []entities.Role) RoleSystem {
	rs := roleSystem{
		totalPermission: entities.PermissionValue(0),
	}

	for _, role := range roles {
		rs.totalPermission |= role.PermissionValue
	}
	return rs
}
