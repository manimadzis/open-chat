package role_service

import (
	"context"
	"open-chat/internal/entities"
)

type RoleService interface {
	Create(ctx context.Context, role *entities.Role, userId entities.UserId, serverId entities.ServerId) error
	Delete(ctx context.Context, roleId entities.RoleId, userId entities.UserId, serverId entities.ServerId) error
	Change(ctx context.Context, role *entities.Role, userId entities.UserId, serverId entities.ServerId) error
	FindByServer(ctx context.Context, serverId entities.ServerId, userId entities.UserId) ([]entities.Role, error)
}
