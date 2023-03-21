package auth_service

import (
	"context"
	"open-chat/internal/entities"
)

type SessionService interface {
	SignUp(ctx context.Context, user entities.User) (*entities.Session, error)
	SignIn(ctx context.Context, user entities.User) (*entities.Session, error)
	LogOut(ctx context.Context, sessionToken entities.SessionToken) error
	FindSessionByToken(ctx context.Context, sessionToken entities.SessionToken) (entities.Session, error)
}
