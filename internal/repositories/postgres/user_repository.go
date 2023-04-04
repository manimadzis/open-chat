package postgres

//
//import (
//	"context"
//	"github.com/jackc/pgx"
//	"open-chat/internal/entities"
//	"open-chat/internal/repositories/user_repository"
//)
//
//type userRepository struct {
//	pool *pgx.ConnPool
//}
//
//func (u userRepository) Create(ctx context.Context, user *entities.User) error {
//	//TODO implement me
//	panic("implement me")
//	//sql := "INSERT INTO user()"
//}
//
//func (u userRepository) FindUserById(ctx context.Context, user *entities.User) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func NewUserRepository(pool *pgx.ConnPool) user_repository.UserRepository {
//	return &userRepository{pool: pool}
//}
