package role_system

import "open-chat/internal/entities"

type roleSystem struct {
	roles           []entities.Role
	totalPermission entities.Permission
}

func (r roleSystem) Check(permissions ...entities.Permission) error {
	sumPermission := entities.Permission(0)
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
		roles:           roles,
		totalPermission: entities.Permission(0),
	}

	for _, role := range roles {
		rs.totalPermission |= role.Permission
	}
	return rs
}
