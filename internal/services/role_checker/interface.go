package role_checker

import (
	"context"
	"open-chat/internal/entities"
)

type RoleChecker interface {
	Check(ctx context.Context, userId entities.UserId, serverId entities.ServerId, permissions ...entities.PermissionValue) error
}
