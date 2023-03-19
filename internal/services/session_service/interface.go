package session_service

import (
	"context"
	"open-chat/internal/entities"
)

type SessionService interface {
	SignUp(ctx context.Context, user *entities.User) (*entities.Session, error)
	SignIn(ctx context.Context, user *entities.User) (*entities.Session, error)
	LogOut(ctx context.Context, session *entities.Session) error
	CheckCredentials(ctx context.Context, session *entities.Session) error
}
