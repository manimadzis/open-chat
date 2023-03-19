package role_system

import "open-chat/internal/entities"

type RoleSystem interface {
	Check(permission ...entities.Permission) error
}
