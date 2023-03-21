package auth_service

import (
	"context"
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"open-chat/internal/entities"
	"open-chat/internal/repositories"
)

type sessionService struct {
	sessionRepo repositories.SessionRepository
	userRepo    repositories.UserRepository
}

func (s *sessionService) createSession(ctx context.Context, userId entities.UserId) (*entities.Session, error) {
	session := entities.NewInfinitySession(generateToken(), userId)
	if err := s.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) SignUp(ctx context.Context, user entities.User) (*entities.Session, error) {
	user.Password = hashPassword(user.Password, user.Login)
	if err := s.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}
	return s.createSession(ctx, user.Id)
}

func (s *sessionService) SignIn(ctx context.Context, user entities.User) (*entities.Session, error) {
	user.Password = hashPassword(user.Password, user.Login)
	systemUser, err := s.userRepo.FindByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}
	if user.Login != systemUser.Login && user.Password != systemUser.Password {
		return nil, ErrInvalidCredentials
	}

	return s.createSession(ctx, systemUser.Id)
}

func (s *sessionService) LogOut(ctx context.Context, sessionToken entities.SessionToken) error {
	return s.sessionRepo.DeleteByToken(ctx, sessionToken)
}

func (s *sessionService) FindSessionByToken(ctx context.Context, sessionToken entities.SessionToken) (entities.Session, error) {
	return s.sessionRepo.FindByToken(ctx, sessionToken)
}

func NewSessionService(userRepo repositories.UserRepository, sessionRepo repositories.SessionRepository) SessionService {
	return &sessionService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func hashPassword(password string, salt string) string {
	return base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 1, 32))
}

func generateToken() entities.SessionToken {
	return entities.SessionToken(uuid.New())
}
