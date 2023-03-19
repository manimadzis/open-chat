package session_service

import (
	"context"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/session_repository"
	"open-chat/internal/repositories/user_repository"
)

type sessionService struct {
	sessionRepo session_repository.SessionRepository
	userRepo    user_repository.UserRepository
}

func (s *sessionService) signIn(ctx context.Context, user *entities.User) (*entities.Session, error) {
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	session := &entities.Session{User: user}
	err := s.sessionRepo.FindByToken(ctx, session)
	if err == nil {
		return session, nil
	}

	session = entities.NewSession(user)
	if err = s.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionService) SignUp(ctx context.Context, user *entities.User) (*entities.Session, error) {
	user.Password = hashPassword(user.Password, user.Login)
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return s.signIn(ctx, user)
}

func (s *sessionService) SignIn(ctx context.Context, user *entities.User) (*entities.Session, error) {
	user.Password = hashPassword(user.Password, user.Login)
	return s.signIn(ctx, user)
}

func (s *sessionService) LogOut(ctx context.Context, session *entities.Session) error {
	return s.sessionRepo.DeleteByToken(ctx, session)
}

func (s *sessionService) CheckCredentials(ctx context.Context, session *entities.Session) error {
	return s.sessionRepo.FindByToken(ctx, session)
}

func NewSessionService(userRepo user_repository.UserRepository, sessionRepo session_repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func hashPassword(password string, salt string) string {
	return base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 1, 32))
}
