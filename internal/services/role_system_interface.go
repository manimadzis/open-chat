package services

import (
	"open-chat/internal/entities"
)

type RoleSystem interface {
	// Check return ErrNotEnoughPermissions if required permissions are not guaranteed by roles.
	Check(permission ...entities.PermissionValue) error

	SetRoles([]entities.Role)
}
