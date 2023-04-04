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
	serverId := entities.ServerId(124)
	server := entities.Server{Id: serverId}

	serverProfileChecker := smock.NewServerProfileChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	serverService := server_service.NewServerService(serverRepo, serverProfileChecker)

	t.Run("cant create in repo",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil
			e := errors.New("some error")
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
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil
			serverRepo.
				On("Create", ctx, mock.AnythingOfType("entities.Server")).
				Return(server.Id, nil).
				Once()
			serverRepo.
				On("Join", ctx, server.Id, server.OwnerId).
				Return(nil)

			servId, err := serverService.Create(ctx, server)
			require.Equal(t, nil, err)
			require.Equal(t, serverId, servId)
		},
	)

}

func TestServerService_Delete(t *testing.T) {
	ctx := context.Background()
	serverId := entities.ServerId(124)
	userId := entities.UserId(444)

	serverProfileChecker := smock.NewServerProfileChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	serverService := server_service.NewServerService(serverRepo, serverProfileChecker)

	t.Run("No permission",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_DELETE_SERVER).
				Return(services.ErrNotEnoughPermissions)

			err := serverService.Delete(ctx, serverId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("Have permission",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
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

	serverProfileChecker := smock.NewServerProfileChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	serverService := server_service.NewServerService(serverRepo, serverProfileChecker)

	t.Run("No permission",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_INVITE_USER).
				Return(services.ErrNotEnoughPermissions)

			err := serverService.Join(ctx, serverId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("Have permission",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_INVITE_USER).
				Return(nil)
			serverRepo.
				On("Join", ctx, serverId, userId).
				Return(nil)

			err := serverService.Join(ctx, serverId, userId)
			require.Equal(t, nil, err)
		},
	)
}

func TestServerService_Kick(t *testing.T) {
	ctx := context.Background()
	serverId := entities.ServerId(124)
	userId := entities.UserId(444)

	serverProfileChecker := smock.NewServerProfileChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	serverService := server_service.NewServerService(serverRepo, serverProfileChecker)

	t.Run("No permission",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_KICK_MEMBER).
				Return(services.ErrNotEnoughPermissions)

			err := serverService.Kick(ctx, serverId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("Have permission",
		func(t *testing.T) {
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
				On("Check", ctx, userId, serverId, role_system.PERM_KICK_MEMBER).
				Return(nil)
			serverRepo.
				On("Kick", ctx, serverId, userId).
				Return(nil)

			err := serverService.Kick(ctx, serverId, userId)
			require.Equal(t, nil, err)
		},
	)
}
