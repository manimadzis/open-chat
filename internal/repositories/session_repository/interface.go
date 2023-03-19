package session_repository

import (
	"context"
	"open-chat/internal/entities"
)

type SessionRepository interface {
	// Create write session. If token already has used it returns ErrTokenAlreadyExists
	Create(ctx context.Context, session *entities.Session) error

	// DeleteByToken delete session by token. If no session found by token it returns ErrNoSuchToken
	DeleteByToken(ctx context.Context, session *entities.Session) error

	// FindByToken find session by token. If no session found by token it returns ErrNoSuchToken
	FindByToken(ctx context.Context, session *entities.Session) error

	// FindByUser find session by user id. If no session found for user it returns ErrUserDoesntHaveSession
	FindByUser(ctx context.Context, session *entities.Session) error
}
