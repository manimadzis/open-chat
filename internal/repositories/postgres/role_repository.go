package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type roleRepository struct {
	pool *pgx.ConnPool
}

func (r roleRepository) FindRolesByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Role, error) {
	sql := `SELECT 
				id,
				name,
				permission,
				creation_time,
				server_id,
				created_by
			FROM role
			WHERE server_id = $1`
	rows, err := r.pool.QueryEx(ctx, sql, nil,
		serverId,
	)
	if err != nil {
		return nil, services.NewUnknownError(err)
	}

	roles := make([]entities.Role, 0)
	for rows.Next() {
		var role entities.Role
		if err := rows.Scan(&role.Id,
			&role.Name,
			&role.PermissionValue,
			&role.CreatedAt,
			&role.ServerId,
			&role.CreatorId,
		); err != nil {
			return nil, services.NewUnknownError(err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r roleRepository) FindPermissionsByValue(ctx context.Context,
	permissionValues []entities.PermissionValue,
) ([]entities.Permission, error) {
	sql := `SELECT value, name, description
			FROM permission
			WHERE value IN $1`
	rows, err := r.pool.QueryEx(ctx, sql, nil,
		permissionValues,
	)
	if err != nil {
		return nil, services.NewUnknownError(err)
	}

	permissions := make([]entities.Permission, 0)
	for rows.Next() {
		var permission entities.Permission
		err := rows.Scan(&permission.Value, &permission.Name, &permission.Description)
		if err != nil {
			return nil, services.NewUnknownError(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func (r roleRepository) Create(ctx context.Context, role entities.Role) (entities.RoleId, error) {
	sql := "INSERT INTO role(name, \"permission\", created_at, server_id, created_by) VALUES($1, $2, $3, $4, $5)"
	row := r.pool.QueryRowEx(ctx, sql, nil,
		role.Name,
		role.PermissionValue,
		role.CreatedAt,
		role.ServerId,
		role.CreatorId,
	)

	if err := row.Scan(&role.Id); err != nil {
		return 0, services.NewUnknownError(err)
	}
	return role.Id, nil
}

func (r roleRepository) Delete(ctx context.Context, roleId entities.RoleId) error {
	sql := "DELETE FROM role WHERE id = $1"
	_, err := r.pool.ExecEx(ctx, sql, nil,
		roleId,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}
	return nil
}

func (r roleRepository) Change(ctx context.Context, role entities.Role) error {
	sql := `UPDATE role 
			SET name = $1,
			permission = $2
			WHERE id = $3`
	_, err := r.pool.ExecEx(ctx, sql, nil,
		role.Name,
		role.PermissionValue,
		role.Id,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}
	return nil
}

func (r roleRepository) FindRoleByServer(ctx context.Context, server *entities.Server) ([]entities.Role, error) {
	sql := `SELECT 
				name, permission, created_at, created_by
			FROM role
			WHERE server_id = $1`
	rows, err := r.pool.QueryEx(ctx, sql, nil, server.Id)

	roles := make([]entities.Role, 0)

	for rows.Next() {
		var role entities.Role
		if err = rows.Scan(
			&role.Name,
			&role.PermissionValue,
			&role.CreatedAt,
			&role.CreatorId,
		); err != nil {
			return nil, services.NewUnknownError(err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func NewRoleRepository(pool *pgx.ConnPool) services.RoleRepository {
	return &roleRepository{
		pool: pool,
	}
}
