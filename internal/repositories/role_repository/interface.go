package role_repository

import (
	"context"
	"open-chat/internal/entities"
)

type RoleRepository interface {
	Create(ctx context.Context, role *entities.Role, user *entities.User) error
	Delete(ctx context.Context, role *entities.Role, user *entities.User) error
	Change(ctx context.Context, role *entities.Role, user *entities.User) error
	FindByServer(ctx context.Context, server *entities.Server, user *entities.User) ([]*entities.Role, error)
}
