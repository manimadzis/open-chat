package channel_service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	rmocks "open-chat/internal/mocks/repositories"
	smocks "open-chat/internal/mocks/services"
	"open-chat/internal/services"
	"open-chat/internal/services/channel_service"
	"open-chat/internal/services/role_system"
	"testing"
)

func TestChannelService_Create(t *testing.T) {
	channelRepo := rmocks.NewChannelRepository(t)
	serverProfileChecker := smocks.NewServerProfileChecker(t)
	serverRepo := rmocks.NewServerRepository(t)

	ctx := context.Background()
	channel := entities.Channel{
		Id: 12,
	}

	t.Run(
		"successfully create channel", func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			channelRepo.
				On("Create", ctx, channel).
				Return(channel.Id, nil).
				Once()
			serverProfileChecker.
				On(
					"Check",
					ctx,
					channel.CreatorId,
					channel.ServerId,
					role_system.PERM_CREATE_CHANNEL,
				).
				Return(nil)
			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			_, err := channelService.Create(ctx, channel)
			require.Equal(t, nil, err)
		},
	)

	t.Run(
		"not enough permissions", func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverProfileChecker.
				On(
					"Check",
					ctx,
					channel.CreatorId,
					channel.ServerId,
					role_system.PERM_CREATE_CHANNEL,
				).
				Return(services.ErrNotEnoughPermissions)

			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			_, err := channelService.Create(ctx, channel)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run(
		"channel repository failed", func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			expectedError := services.UnknownError(errors.New("some happened"))
			serverProfileChecker.
				On(
					"Check",
					ctx,
					channel.CreatorId,
					channel.ServerId,
					role_system.PERM_CREATE_CHANNEL,
				).
				Return(nil)

			channelRepo.
				On("Create", ctx, channel).
				Return(entities.ChannelId(0), expectedError)

			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			_, err := channelService.Create(ctx, channel)
			require.Equal(t, expectedError, err)
		},
	)
}

func TestChannelService_Delete(t *testing.T) {
	channelRepo := rmocks.NewChannelRepository(t)
	serverProfileChecker := smocks.NewServerProfileChecker(t)
	serverRepo := rmocks.NewServerRepository(t)

	ctx := context.Background()
	channelId := entities.ChannelId(123)
	userId := entities.UserId(423)
	serverId := entities.ServerId(65)
	server := &entities.Server{Id: serverId}

	t.Run("successfully delete channel",
		func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			channelRepo.
				On("Delete", ctx, channelId).
				Return(nil).
				Once()

			serverProfileChecker.
				On(
					"Check",
					ctx,
					userId,
					serverId,
					role_system.PERM_DELETE_CHANNEL,
				).
				Return(nil)

			serverRepo.
				On("FindByChannelId", ctx, channelId).
				Return(server, nil)

			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			err := channelService.Delete(ctx, channelId, userId)
			require.Equal(t, nil, err)
		},
	)

	t.Run("not enough permissions",
		func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverRepo.
				On("FindByChannelId", ctx, channelId).
				Return(server, nil)

			serverProfileChecker.
				On(
					"Check",
					ctx,
					userId,
					serverId,
					role_system.PERM_DELETE_CHANNEL,
				).
				Return(services.ErrNotEnoughPermissions)

			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			err := channelService.Delete(ctx, channelId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("no such channel",
		func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			serverRepo.
				On("FindByChannelId", ctx, channelId).
				Return(server, services.ErrNoSuchChannel)

			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			err := channelService.Delete(ctx, channelId, userId)
			require.Equal(t, services.ErrNoSuchChannel, err)
		},
	)
}

func TestChannelService_FindByServerId(t *testing.T) {
	channelRepo := rmocks.NewChannelRepository(t)
	serverProfileChecker := smocks.NewServerProfileChecker(t)
	serverRepo := rmocks.NewServerRepository(t)

	ctx := context.Background()
	serverId := entities.ServerId(123)
	userId := entities.UserId(1235)

	t.Run("found",
		func(t *testing.T) {
			channelRepo.ExpectedCalls = nil
			serverProfileChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil

			channelRepo.
				On("FindByServerId", ctx, serverId).
				Return(nil, nil).
				Once()

			serverProfileChecker.
				On(
					"Check",
					ctx,
					userId,
					serverId,
				).
				Return(nil)
			channelService := channel_service.NewChannelService(
				channelRepo,
				serverProfileChecker,
				serverRepo,
			)

			_, err := channelService.FindByServerId(ctx, serverId, userId)
			require.Equal(t, nil, err)
		},
	)

	t.Run("user doesn't have server profile", func(t *testing.T) {
		channelRepo.ExpectedCalls = nil
		serverProfileChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil

		serverProfileChecker.
			On(
				"Check",
				ctx,
				userId,
				serverId,
			).
			Return(services.ErrNoSuchServerProfile)

		channelService := channel_service.NewChannelService(
			channelRepo,
			serverProfileChecker,
			serverRepo,
		)

		_, err := channelService.FindByServerId(ctx, serverId, userId)
		require.Equal(t, services.ErrNoSuchServerProfile, err)
	},
	)
}
