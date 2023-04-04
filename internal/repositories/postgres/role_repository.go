package postgres

//
//import (
//	"context"
//	"errors"
//	"github.com/jackc/pgx"
//	"log"
//	"open-chat/internal/entities"
//	"open-chat/internal/repositories"
//)
//
//type roleRepository struct {
//	pool *pgx.ConnPool
//}
//
//func (r roleRepository) FindRolesByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Role, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r roleRepository) FindPermissionsByValue(ctx context.Context, permission []entities.PermissionValue) ([]entities.Permission, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//// TODO: перепроверить логику связанную с обработкой ошибок
//
//func (r roleRepository) Create(ctx context.Context, role *entities.Role) error {
//	tx, err := r.pool.BeginEx(ctx, nil)
//	if err != nil {
//		return repositories.UnknownError(err)
//	}
//	defer func() {
//		switch err {
//		case nil:
//			err = tx.Commit()
//		default:
//			_ = tx.Rollback()
//		}
//	}()
//
//	sql := "INSERT INTO role(name, \"permission\", created_at, server_id, created_by) VALUES($1, $2, $3, $4, $5)"
//	_, err = tx.Exec(sql, role.Name, role.Permission, role.CreatedAt, role.Server.Id, role.CreatedBy.Id)
//	if err != nil {
//		var pgErr pgx.PgError
//		if errors.As(err, &pgErr) {
//			log.Print(pgErr)
//		}
//		return err
//	}
//	return err
//}
//
//func (r roleRepository) Delete(ctx context.Context, roleId entities.RoleId) error {
//	tx, err := r.pool.BeginEx(ctx, nil)
//	if err != nil {
//		return repositories.UnknownError(err)
//	}
//	defer func() {
//		switch err {
//		case nil:
//			err = tx.Commit()
//		default:
//			_ = tx.Rollback()
//		}
//	}()
//
//	sql := "DELETE FROM role WHERE id = $1"
//	_, err = tx.Exec(sql, roleId)
//	if err != nil {
//		var pgErr pgx.PgError
//		if errors.As(err, &pgErr) {
//			log.Print(pgErr)
//		}
//		return err
//	}
//
//	return err
//}
//
//func (r roleRepository) Change(ctx context.Context, role *entities.Role) error {
//	panic("no impl")
//}
//
//func (r roleRepository) FindRoleByServer(ctx context.Context, server *entities.Server) ([]*entities.Role, error) {
//	tx, err := r.pool.BeginEx(ctx, nil)
//	if err != nil {
//		return nil, repositories.UnknownError(err)
//	}
//	defer func() {
//		switch err {
//		case nil:
//			err = tx.Commit()
//		default:
//			_ = tx.Rollback()
//		}
//	}()
//
//	sql := "SELECT name, permission, created_at, created_by FROM role WHERE server_id = $1"
//	rows, err := tx.Query(sql, server.Id)
//	if err != nil {
//		var pgErr pgx.PgError
//		if errors.As(err, &pgErr) {
//			log.Print(pgErr)
//		}
//		return nil, err
//	}
//
//	roles := make([]*entities.Role, 0)
//
//	for rows.Next() {
//		role := &entities.Role{CreatedBy: &entities.User{}}
//		err = rows.Scan(&role.Name, &role.Permission, &role.CreatedAt, &role.CreatedBy.Id)
//		if err != nil {
//			return nil, err
//		}
//		roles = append(roles, role)
//	}
//
//	return roles, err
//}
//
//func NewRoleRepository(pool *pgx.ConnPool) repositories.RoleRepository {
//	return &roleRepository{
//		pool: pool,
//	}
//}
