package auth_service

import (
	"context"
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type authService struct {
	sessionRepo services.SessionRepository
	userRepo    services.UserRepository
}

func (a *authService) createSession(ctx context.Context, userId entities.UserId) (*entities.Session, error) {
	session := entities.NewInfinitySession(generateToken(), userId)
	if err := a.sessionRepo.Create(ctx, *session); err != nil {
		return nil, err
	}
	return session, nil
}

func (a *authService) SignUp(ctx context.Context, user entities.User) (*entities.Session, error) {
	var err error
	user.Password = hashPassword(user.Password, user.Login)
	user.Id, err = a.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return a.createSession(ctx, user.Id)
}

func (a *authService) SignIn(ctx context.Context, user entities.User) (*entities.Session, error) {
	user.Password = hashPassword(user.Password, user.Login)
	systemUser, err := a.userRepo.FindByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}
	if user.Password != systemUser.Password {
		return nil, services.ErrInvalidCredentials
	}

	return a.createSession(ctx, systemUser.Id)
}

func (a *authService) LogOut(ctx context.Context, sessionToken entities.SessionToken) error {
	return a.sessionRepo.DeleteByToken(ctx, sessionToken)
}

func (a *authService) FindSessionByToken(ctx context.Context, sessionToken entities.SessionToken) (*entities.Session,
	error,
) {
	return a.sessionRepo.FindByToken(ctx, sessionToken)
}

func hashPassword(password string, salt string) string {
	return base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 1, 32))
}

func generateToken() entities.SessionToken {
	return entities.SessionToken(uuid.New())
}

func NewAuthService(sessionRepo services.SessionRepository, userRepo services.UserRepository) services.AuthService {
	return &authService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}
