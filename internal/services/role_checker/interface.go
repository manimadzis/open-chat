package role_checker

import (
	"context"
	"open-chat/internal/entities"
)

type RoleChecker interface {
	Check(ctx context.Context, user *entities.User, server *entities.Server, permissions ...entities.Permission) error
}
