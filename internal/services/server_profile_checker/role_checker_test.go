package server_profile_checker

import (
	"context"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	rmock "open-chat/internal/mocks/repositories"
	smock "open-chat/internal/mocks/services"
	"open-chat/internal/services"
	"testing"
)

func TestRoleChecker_Check(t *testing.T) {
	serverRepo := rmock.NewServerRepository(t)
	roleSystem := smock.NewRoleSystem(t)

	ctx := context.Background()
	userId := entities.UserId(123)
	serverId := entities.ServerId(1234)
	serverProfile := &entities.ServerProfile{}

	t.Run("User doesn't have server profile",
		func(t *testing.T) {
			serverRepo.ExpectedCalls = nil
			roleSystem.ExpectedCalls = nil
			serverProfileChecker := NewServerProfileChecker(serverRepo, roleSystem)

			serverRepo.
				On("FindServerProfileByIds", ctx, serverId, userId).
				Return(serverProfile, services.ErrNoSuchServerProfile)

			err := serverProfileChecker.Check(ctx, userId, serverId)
			require.Equal(t, services.ErrNoSuchServerProfile, err)
		},
	)

	t.Run("User doesn't have enough permissions",
		func(t *testing.T) {
			serverRepo.ExpectedCalls = nil
			roleSystem.ExpectedCalls = nil

			serverProfileChecker := NewServerProfileChecker(serverRepo, roleSystem)
			serverRepo.
				On("FindServerProfileByIds", ctx, serverId, userId).
				Return(serverProfile, nil)

			roleSystem.
				On("SetRoles", serverProfile.Roles).
				Return().
				Once()

			roleSystem.
				On("Check").
				Return(services.ErrNotEnoughPermissions).
				Once()

			err := serverProfileChecker.Check(ctx, userId, serverId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("User have enough permissions",
		func(t *testing.T) {
			serverRepo.ExpectedCalls = nil
			roleSystem.ExpectedCalls = nil

			serverProfileChecker := NewServerProfileChecker(serverRepo, roleSystem)
			serverRepo.
				On("FindServerProfileByIds", ctx, serverId, userId).
				Return(serverProfile, nil)

			roleSystem.
				On("SetRoles", serverProfile.Roles).
				Return().
				Once()

			roleSystem.
				On("Check").
				Return(nil).
				Once()

			err := serverProfileChecker.Check(ctx, userId, serverId)
			require.Equal(t, nil, err)
		},
	)
}
