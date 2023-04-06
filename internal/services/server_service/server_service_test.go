package server_service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	rmock "open-chat/internal/mocks/repositories"
	smock "open-chat/internal/mocks/services"
	"open-chat/internal/services"
	"open-chat/internal/services/role_system"
	"open-chat/internal/services/server_service"
	"testing"
)

func TestServerService_Create(t *testing.T) {
	ctx := context.Background()
	server := entities.Server{Id: 123, OwnerId: 213}
	user := entities.User{}

	permissionChecker := smock.NewPermissionChecker(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)
	serverRepo := rmock.NewServerRepository(t)
	userRepo := rmock.NewUserRepository(t)

	clearMocks := func() {
		permissionChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil
		serverProfileRepo.ExpectedCalls = nil
	}

	serverService := server_service.NewServerService(serverRepo, serverProfileRepo, userRepo, permissionChecker)

	t.Run("cant find owner",
		func(t *testing.T) {
			clearMocks()
			e := errors.New("some error")
			userRepo.
				On("FindById", ctx, server.OwnerId).
				Return(nil, e).
				Once()

			_, err := serverService.Create(ctx, server)
			require.Equal(t, e, err)
		},
	)

	t.Run("cant create in repo",
		func(t *testing.T) {
			permissionChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil
			e := errors.New("some error")

			userRepo.
				On("FindById", ctx, server.OwnerId).
				Return(&user, nil).
				Once()

			serverRepo.
				On("Create", ctx, mock.AnythingOfType("entities.Server")).
				Return(server.Id, e).
				Once()

			_, err := serverService.Create(ctx, server)
			require.Equal(t, e, err)
		},
	)

	t.Run("success",
		func(t *testing.T) {
			clearMocks()
			userRepo.
				On("FindById", ctx, server.OwnerId).
				Return(&user, nil).
				Once()

			serverRepo.
				On("Create", ctx, mock.AnythingOfType("entities.Server")).
				Return(server.Id, nil).
				Once()

			serverProfileRepo.
				On("Create", ctx, mock.AnythingOfType("entities.ServerProfile")).
				Return(entities.ServerProfileId{}, nil).
				Once()

			servId, err := serverService.Create(ctx, server)
			require.Equal(t, nil, err)
			require.Equal(t, server.Id, servId)
		},
	)

}

func TestServerService_Delete(t *testing.T) {
	ctx := context.Background()
	serverId := entities.ServerId(124)
	userId := entities.UserId(444)

	permissionChecker := smock.NewPermissionChecker(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)
	serverRepo := rmock.NewServerRepository(t)
	userRepo := rmock.NewUserRepository(t)

	serverService := server_service.NewServerService(serverRepo, serverProfileRepo, userRepo, permissionChecker)

	t.Run("No permission",
		func(t *testing.T) {
			permissionChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_DELETE_SERVER).
				Return(services.ErrNotEnoughPermissions)

			err := serverService.Delete(ctx, serverId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("Have permission",
		func(t *testing.T) {
			permissionChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_DELETE_SERVER).
				Return(nil)
			serverRepo.
				On("Delete", ctx, serverId).
				Return(nil)

			err := serverService.Delete(ctx, serverId, userId)
			require.Equal(t, nil, err)
		},
	)
}

func TestServerService_Join(t *testing.T) {
	ctx := context.Background()
	serverId := entities.ServerId(124)
	userId := entities.UserId(444)
	user := entities.User{
		Id:       userId,
		Nickname: "123",
	}

	permissionChecker := smock.NewPermissionChecker(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)
	serverRepo := rmock.NewServerRepository(t)
	userRepo := rmock.NewUserRepository(t)

	clearMocks := func() {
		permissionChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil
		serverProfileRepo.ExpectedCalls = nil
	}

	serverService := server_service.NewServerService(serverRepo, serverProfileRepo, userRepo, permissionChecker)

	t.Run("No permission",
		func(t *testing.T) {
			clearMocks()
			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_INVITE_USER).
				Return(services.ErrNotEnoughPermissions).
				Once()

			err := serverService.Join(ctx, serverId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("Have permission and don't find user",
		func(t *testing.T) {
			clearMocks()
			e := errors.New("some error")
			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_INVITE_USER).
				Return(nil).
				Once()
			userRepo.
				On("FindById", ctx, userId).
				Return(&user, e).
				Once()

			err := serverService.Join(ctx, serverId, userId)
			require.Equal(t, e, err)
		},
	)

	t.Run("Have permission and find user",
		func(t *testing.T) {
			clearMocks()
			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_INVITE_USER).
				Return(nil).
				Once()
			userRepo.
				On("FindById", ctx, userId).
				Return(&user, nil).
				Once()
			serverProfileRepo.
				On("Create", ctx, mock.AnythingOfType("entities.ServerProfile")).
				Return(entities.ServerProfileId{}, nil).
				Once()

			err := serverService.Join(ctx, serverId, userId)
			require.Equal(t, nil, err)
		},
	)
}

func TestServerService_Kick(t *testing.T) {
	ctx := context.Background()
	serverId := entities.ServerId(124)
	userId := entities.UserId(444)

	permissionChecker := smock.NewPermissionChecker(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)
	serverRepo := rmock.NewServerRepository(t)
	userRepo := rmock.NewUserRepository(t)

	clearMocks := func() {
		permissionChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil
		userRepo.ExpectedCalls = nil
		serverProfileRepo.ExpectedCalls = nil
	}

	serverService := server_service.NewServerService(serverRepo, serverProfileRepo, userRepo, permissionChecker)

	t.Run("No permission",
		func(t *testing.T) {
			clearMocks()
			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_KICK_MEMBER).
				Return(services.ErrNotEnoughPermissions)

			err := serverService.Kick(ctx, serverId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("Have permission",
		func(t *testing.T) {
			clearMocks()
			permissionChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_KICK_MEMBER).
				Return(nil)
			serverProfileRepo.
				On("Delete", ctx, entities.ServerProfileId{
					UserId:   userId,
					ServerId: serverId,
				},
				).
				Return(nil)

			err := serverService.Kick(ctx, serverId, userId)
			require.Equal(t, nil, err)
		},
	)
}
