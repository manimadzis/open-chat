package auth_service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	rmocks "open-chat/internal/mocks/repositories"
	"open-chat/internal/services"
	"testing"
)

func TestAuthService_LogOut(t *testing.T) {
	sessionRepo := rmocks.NewSessionRepository(t)
	userRepo := rmocks.NewUserRepository(t)
	ctx := context.Background()
	sessionToken := entities.SessionToken{}

	t.Run("success logout", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil
		sessionRepo.
			On("DeleteByToken", ctx, sessionToken).
			Return(nil).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		err := sessionService.LogOut(ctx, sessionToken)
		require.Equal(t, nil, err)
	},
	)

	t.Run("failed logout", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		sessionRepo.
			On("DeleteByToken", ctx, sessionToken).
			Return(services.ErrNoSuchToken).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		err := sessionService.LogOut(ctx, sessionToken)
		require.Equal(t, services.ErrNoSuchToken, err)
	},
	)
}

func TestAuthService_FindSessionByToken(t *testing.T) {
	sessionRepo := rmocks.NewSessionRepository(t)
	userRepo := rmocks.NewUserRepository(t)
	ctx := context.Background()
	sessionToken := entities.SessionToken{}
	session := &entities.Session{UserId: 123}

	t.Run("success", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil
		sessionRepo.
			On("FindByToken", ctx, sessionToken).
			Return(session, nil).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		retSession, err := sessionService.FindSessionByToken(ctx, sessionToken)
		require.Equal(t, nil, err)
		require.Equal(t, session, retSession)
	},
	)

	t.Run("no such token", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		sessionRepo.
			On("FindByToken", ctx, sessionToken).
			Return(session, services.ErrNoSuchToken).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		retSession, err := sessionService.FindSessionByToken(ctx, sessionToken)
		require.Equal(t, services.ErrNoSuchToken, err)
		require.Equal(t, session, retSession)
	},
	)
}

func TestAuthService_SignIn(t *testing.T) {
	sessionRepo := rmocks.NewSessionRepository(t)
	userRepo := rmocks.NewUserRepository(t)

	ctx := context.Background()
	user := entities.User{
		Login:    "123",
		Password: "123",
	}
	hashedUser := user
	hashedUser.Password = hashPassword(hashedUser.Password, hashedUser.Login)

	t.Run("success sign in", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil

		userRepo.
			On("FindByLogin", ctx, user.Login).
			Return(&hashedUser, nil)

		sessionRepo.
			On("Create", ctx, mock.AnythingOfType("entities.Session")).
			Return(nil).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		_, err := sessionService.SignIn(ctx, user)
		require.Equal(t, nil, err)
	},
	)

	t.Run("invalid login", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil

		userRepo.
			On("FindByLogin", ctx, user.Login).
			Return(nil, services.ErrNoSuchLogin).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		_, err := sessionService.SignIn(ctx, user)
		require.Equal(t, services.ErrNoSuchLogin, err)
	},
	)

	t.Run("invalid password", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil

		hashedUser.Password = "wrong pass"

		userRepo.
			On("FindByLogin", ctx, user.Login).
			Return(&hashedUser, nil).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		_, err := sessionService.SignIn(ctx, user)
		require.Equal(t, services.ErrInvalidCredentials, err)
	},
	)
}

func TestAuthService_SignUp(t *testing.T) {
	sessionRepo := rmocks.NewSessionRepository(t)
	userRepo := rmocks.NewUserRepository(t)

	ctx := context.Background()
	user := entities.User{
		Id:       123,
		Login:    "123",
		Password: "123",
	}
	hashedUser := user
	hashedUser.Password = hashPassword(hashedUser.Password, hashedUser.Login)

	t.Run("success sign up", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil

		userRepo.
			On("Create", ctx, hashedUser).
			Return(hashedUser.Id, nil).
			Once()

		sessionRepo.
			On("Create", ctx, mock.AnythingOfType("entities.Session")).
			Return(nil).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		_, err := sessionService.SignUp(ctx, user)
		require.Equal(t, nil, err)
	},
	)

	t.Run("already used login", func(t *testing.T) {
		sessionRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil

		userRepo.
			On("Create", ctx, hashedUser).
			Return(hashedUser.Id, services.ErrLoginAlreadyExists).
			Once()

		sessionService := NewAuthService(sessionRepo, userRepo)

		_, err := sessionService.SignUp(ctx, user)
		require.Equal(t, services.ErrLoginAlreadyExists, err)
	},
	)
}
