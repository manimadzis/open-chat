package role_service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	rmock "open-chat/internal/mocks/repositories"
	smock "open-chat/internal/mocks/services"
	"open-chat/internal/services"
	"open-chat/internal/services/role_service"
	"open-chat/internal/services/role_system"
	"testing"
)

func TestRoleService_Create(t *testing.T) {
	ctx := context.Background()
	role := entities.Role{
		ServerId:  123,
		CreatorId: 345,
	}
	roleId := entities.RoleId(543)

	serverProfileChecker := smock.NewServerProfileChecker(t)
	roleRepo := rmock.NewRoleRepository(t)
	roleService := role_service.NewRoleService(roleRepo, serverProfileChecker)

	t.Run("Enough permissions", func(t *testing.T) {
		{
			serverProfileChecker.ExpectedCalls = nil
			roleRepo.ExpectedCalls = nil

			serverProfileChecker.
				On(
					"Check",
					ctx,
					role.CreatorId,
					role.ServerId,
					role_system.PERM_CREATE_ROLE,
				).
				Return(nil).
				Once()
			roleRepo.
				On("Create", ctx, role).
				Return(roleId, nil).
				Once()

			id, err := roleService.Create(ctx, role)
			require.Equal(t, roleId, id)
			require.Equal(t, nil, err)
		}
		{
			serverProfileChecker.ExpectedCalls = nil
			roleRepo.ExpectedCalls = nil
			expectedErr := services.NewUnknownError(errors.New("some happened"))
			serverProfileChecker.
				On(
					"Check",
					ctx,
					role.CreatorId,
					role.ServerId,
					role_system.PERM_CREATE_ROLE,
				).
				Return(nil).
				Once()
			roleRepo.
				On("Create", ctx, role).
				Return(roleId, expectedErr).
				Once()

			_, err := roleService.Create(ctx, role)
			require.Equal(t, expectedErr, err)
		}
	},
	)

	t.Run(
		"Not enough permissions", func(t *testing.T) {

			{
				serverProfileChecker.ExpectedCalls = nil
				roleRepo.ExpectedCalls = nil

				serverProfileChecker.
					On(
						"Check",
						ctx,
						role.CreatorId,
						role.ServerId,
						role_system.PERM_CREATE_ROLE,
					).
					Return(services.ErrNotEnoughPermissions).
					Once()
				roleRepo.
					On("Create", ctx, role).
					Return(roleId, nil).
					Unset()

				_, err := roleService.Create(ctx, role)
				require.Equal(t, services.ErrNotEnoughPermissions, err)
			}
		},
	)
}

func TestRoleService_Delete(t *testing.T) {
	ctx := context.Background()

	roleId := entities.RoleId(543)
	userId := entities.UserId(333)
	serverId := entities.ServerId(222)

	serverProfileChecker := smock.NewServerProfileChecker(t)
	roleRepo := rmock.NewRoleRepository(t)
	roleService := role_service.NewRoleService(roleRepo, serverProfileChecker)

	t.Run(
		"Enough permissions", func(t *testing.T) {
			{
				serverProfileChecker.ExpectedCalls = nil
				roleRepo.ExpectedCalls = nil

				serverProfileChecker.
					On(
						"Check",
						ctx,
						userId,
						serverId,
						role_system.PERM_DELETE_ROLE,
					).
					Return(nil).
					Once()
				roleRepo.
					On("Delete", ctx, roleId).
					Return(nil).
					Once()

				err := roleService.Delete(ctx, roleId, userId, serverId)
				require.Equal(t, nil, err)
			}
		},
	)

	t.Run("Not enough permissions", func(t *testing.T) {
		serverProfileChecker.ExpectedCalls = nil
		roleRepo.ExpectedCalls = nil

		serverProfileChecker.
			On(
				"Check",
				ctx,
				userId,
				serverId,
				role_system.PERM_DELETE_ROLE,
			).
			Return(services.ErrNotEnoughPermissions).
			Once()

		err := roleService.Delete(ctx, roleId, userId, serverId)
		require.Equal(t, services.ErrNotEnoughPermissions, err)
	},
	)
}

func TestRoleService_Change(t *testing.T) {
	ctx := context.Background()

	role := entities.Role{
		Id:   123,
		Name: "",
	}
	userId := entities.UserId(333)
	serverId := entities.ServerId(222)

	serverProfileChecker := smock.NewServerProfileChecker(t)
	roleRepo := rmock.NewRoleRepository(t)
	roleService := role_service.NewRoleService(roleRepo, serverProfileChecker)

	t.Run(
		"Enough permissions", func(t *testing.T) {
			{
				serverProfileChecker.ExpectedCalls = nil
				roleRepo.ExpectedCalls = nil

				serverProfileChecker.
					On(
						"Check",
						ctx,
						userId,
						serverId,
						role_system.PERM_CHANGE_ROLE,
					).
					Return(nil).
					Once()
				roleRepo.
					On("Change", ctx, role).
					Return(nil).
					Once()

				err := roleService.Change(ctx, role, userId, serverId)
				require.Equal(t, nil, err)
			}
		},
	)

	t.Run("Not enough permissions", func(t *testing.T) {
		serverProfileChecker.ExpectedCalls = nil
		roleRepo.ExpectedCalls = nil

		serverProfileChecker.
			On(
				"Check",
				ctx,
				userId,
				serverId,
				role_system.PERM_CHANGE_ROLE,
			).
			Return(services.ErrNotEnoughPermissions).
			Once()

		err := roleService.Change(ctx, role, userId, serverId)
		require.Equal(t, services.ErrNotEnoughPermissions, err)
	},
	)
}

func TestRoleService_FindByServer(t *testing.T) {
	ctx := context.Background()

	userId := entities.UserId(333)
	serverId := entities.ServerId(222)

	serverProfileChecker := smock.NewServerProfileChecker(t)
	roleRepo := rmock.NewRoleRepository(t)
	roleService := role_service.NewRoleService(roleRepo, serverProfileChecker)

	t.Run(
		"Enough permissions", func(t *testing.T) {
			{
				serverProfileChecker.ExpectedCalls = nil
				roleRepo.ExpectedCalls = nil

				serverProfileChecker.
					On(
						"Check",
						ctx,
						userId,
						serverId,
					).
					Return(nil).
					Once()
				roleRepo.
					On("FindRolesByServerId", ctx, serverId).
					Return(nil, nil).
					Once()

				roles, err := roleService.FindByServer(ctx, serverId, userId)
				require.Equal(t, nil, err)
				require.Equal(t, []entities.Role(nil), roles)
			}
		},
	)

	t.Run("Not enough permissions", func(t *testing.T) {
		serverProfileChecker.ExpectedCalls = nil
		roleRepo.ExpectedCalls = nil

		serverProfileChecker.
			On(
				"Check",
				ctx,
				userId,
				serverId,
			).
			Return(services.ErrNotEnoughPermissions).
			Once()

		roles, err := roleService.FindByServer(ctx, serverId, userId)
		require.Equal(t, services.ErrNotEnoughPermissions, err)
		require.Equal(t, []entities.Role(nil), roles)
	},
	)
}
