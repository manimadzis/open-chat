package message_service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	rmock "open-chat/internal/mocks/repositories"
	smock "open-chat/internal/mocks/services"
	"open-chat/internal/services"
	"open-chat/internal/services/message_service"
	"open-chat/internal/services/role_system"
	"testing"
	"time"
)

func TestMessageService_Send(t *testing.T) {
	ctx := context.Background()
	message := entities.Message{Text: "Aboba"}
	server := &entities.Server{}
	messageId := entities.MessageId(10)

	permissionChecker := smock.NewPermissionChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	messageRepo := rmock.NewMessageRepository(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)

	clearMocks := func() {
		permissionChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil
		messageRepo.ExpectedCalls = nil
		serverProfileRepo.ExpectedCalls = nil
	}

	messageService := message_service.NewMessageService(messageRepo, serverRepo, serverProfileRepo, permissionChecker)

	t.Run("Enough permissions", func(t *testing.T) {
		clearMocks()

		serverRepo.
			On("FindByChannelId", ctx, message.ChannelId).
			Return(server, nil).
			Once()

		permissionChecker.
			On(
				"Check",
				ctx,
				message.SenderId,
				server.Id,
				role_system.PERM_SEND_MESSAGE,
			).
			Return(nil).
			Once()

		messageRepo.
			On("Create", ctx, message).
			Return(messageId, nil)

		msgId, err := messageService.Send(ctx, message)
		require.Equal(t, nil, err)
		require.Equal(t, messageId, msgId)
	},
	)

	t.Run(
		"Not enough permissions", func(t *testing.T) {
			permissionChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil
			messageRepo.ExpectedCalls = nil

			serverRepo.
				On("FindByChannelId", ctx, message.ChannelId).
				Return(server, nil).
				Once()

			permissionChecker.
				On(
					"Check",
					ctx,
					message.SenderId,
					server.Id,
					role_system.PERM_SEND_MESSAGE,
				).
				Return(services.ErrNotEnoughPermissions).
				Once()

			_, err := messageService.Send(ctx, message)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run(
		"No server profile", func(t *testing.T) {
			permissionChecker.ExpectedCalls = nil
			serverRepo.ExpectedCalls = nil
			messageRepo.ExpectedCalls = nil

			serverRepo.
				On("FindByChannelId", ctx, message.ChannelId).
				Return(nil, services.ErrNoSuchServerProfile).
				Once()

			_, err := messageService.Send(ctx, message)
			require.Equal(t, services.ErrNoSuchServerProfile, err)
		},
	)

}

func TestMessageService_Delete(t *testing.T) {
	ctx := context.Background()
	messageId := entities.MessageId(111)
	userId := entities.UserId(1112)
	server := &entities.Server{}

	permissionChecker := smock.NewPermissionChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	messageRepo := rmock.NewMessageRepository(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)

	clearMocks := func() {
		permissionChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil
		messageRepo.ExpectedCalls = nil
		serverProfileRepo.ExpectedCalls = nil
	}

	messageService := message_service.NewMessageService(messageRepo, serverRepo, serverProfileRepo, permissionChecker)

	t.Run("Enough permissions",
		func(t *testing.T) {
			clearMocks()

			serverRepo.
				On("FindByMessageId", ctx, messageId).
				Return(server, nil).
				Once()

			permissionChecker.
				On(
					"Check",
					ctx,
					userId,
					server.Id,
					role_system.PERM_DELETE_MESSAGE,
				).
				Return(nil).
				Once()

			messageRepo.
				On("Delete", ctx, messageId).
				Return(nil)

			err := messageService.Delete(ctx, messageId, userId)
			require.Equal(t, nil, err)
		},
	)

	t.Run("Not enough permissions",
		func(t *testing.T) {
			clearMocks()

			serverRepo.
				On("FindByMessageId", ctx, messageId).
				Return(server, nil).
				Once()

			permissionChecker.
				On(
					"Check",
					ctx,
					userId,
					server.Id,
					role_system.PERM_DELETE_MESSAGE,
				).
				Return(services.ErrNotEnoughPermissions).
				Once()

			err := messageService.Delete(ctx, messageId, userId)
			require.Equal(t, services.ErrNotEnoughPermissions, err)
		},
	)

	t.Run("No server profile",
		func(t *testing.T) {
			clearMocks()

			serverRepo.
				On("FindByMessageId", ctx, messageId).
				Return(nil, services.ErrNoSuchServerProfile).
				Once()

			err := messageService.Delete(ctx, messageId, userId)
			require.Equal(t, services.ErrNoSuchServerProfile, err)
		},
	)
}

func TestMessageService_FindInChat(t *testing.T) {
	ctx := context.Background()
	userId := entities.UserId(123)
	channelId := entities.ChannelId(42)
	filters := entities.MessageFiltersDTO{}
	server := &entities.Server{}

	permissionChecker := smock.NewPermissionChecker(t)
	serverRepo := rmock.NewServerRepository(t)
	messageRepo := rmock.NewMessageRepository(t)
	serverProfileRepo := rmock.NewServerProfileRepository(t)

	clearMocks := func() {
		permissionChecker.ExpectedCalls = nil
		serverRepo.ExpectedCalls = nil
		messageRepo.ExpectedCalls = nil
		serverProfileRepo.ExpectedCalls = nil
	}

	messageService := message_service.NewMessageService(messageRepo, serverRepo, serverProfileRepo, permissionChecker)

	t.Run("No channel", func(t *testing.T) {
		clearMocks()

		serverRepo.
			On("FindByChannelId", ctx, channelId).
			Return(nil, services.ErrNoSuchChannel).
			Once()

		_, err := messageService.FindInChat(ctx, userId, channelId, filters)
		require.Equal(t, services.ErrNoSuchChannel, err)
	},
	)

	t.Run("No permissions", func(t *testing.T) {
		clearMocks()

		serverRepo.
			On("FindByChannelId", ctx, channelId).
			Return(server, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE).
			Return(services.ErrNotEnoughPermissions).
			Once()

		_, err := messageService.FindInChat(ctx, userId, channelId, filters)
		require.Equal(t, services.ErrNotEnoughPermissions, err)
	},
	)

	t.Run("message repo error", func(t *testing.T) {
		clearMocks()

		e := errors.New("some error")
		serverRepo.
			On("FindByChannelId", ctx, channelId).
			Return(server, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE).
			Return(nil).
			Once()

		messageRepo.
			On("FindByChannelId", ctx, channelId, filters.Offset, filters.Count).
			Return(nil, e).
			Once()

		_, err := messageService.FindInChat(ctx, userId, channelId, filters)
		require.Equal(t, e, err)
	},
	)

	t.Run("has permission for history and success", func(t *testing.T) {
		clearMocks()

		serverRepo.
			On("FindByChannelId", ctx, channelId).
			Return(server, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE).
			Return(nil).
			Once()

		messageRepo.
			On("FindByChannelId", ctx, channelId, filters.Offset, filters.Count).
			Return(nil, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE_HISTORY).
			Return(nil).
			Once()

		_, err := messageService.FindInChat(ctx, userId, channelId, filters)
		require.Equal(t, nil, err)
	},
	)

	t.Run("No permission for history, but success", func(t *testing.T) {
		clearMocks()

		ct := time.Now()

		serverRepo.
			On("FindByChannelId", ctx, channelId).
			Return(server, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE).
			Return(nil).
			Once()

		messageRepo.
			On("FindByChannelId", ctx, channelId, filters.Offset, filters.Count).
			Return([]entities.Message{{Time: ct}, {Time: ct.Add(100 * time.Minute)}}, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE_HISTORY).
			Return(services.ErrNotEnoughPermissions).
			Once()
		serverProfileRepo.
			On("FindById", ctx, entities.ServerProfileId{ServerId: server.Id, UserId: userId}).
			Return(&entities.ServerProfile{JoinTime: ct.Add(50 * time.Minute)}, nil).
			Once()

		_, err := messageService.FindInChat(ctx, userId, channelId, filters)
		require.Equal(t, nil, err)
	},
	)

	t.Run("No permission for history, but failed to get server profile", func(t *testing.T) {
		clearMocks()
		e := errors.New("some error")
		serverRepo.
			On("FindByChannelId", ctx, channelId).
			Return(server, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE).
			Return(nil).
			Once()

		messageRepo.
			On("FindByChannelId", ctx, channelId, filters.Offset, filters.Count).
			Return(nil, nil).
			Once()

		permissionChecker.
			On("Check", ctx, userId, server.Id, role_system.PERM_READ_MESSAGE_HISTORY).
			Return(services.ErrNotEnoughPermissions).
			Once()
		serverProfileRepo.
			On("FindById", ctx, entities.ServerProfileId{ServerId: server.Id, UserId: userId}).
			Return(&entities.ServerProfile{}, e).
			Once()

		_, err := messageService.FindInChat(ctx, userId, channelId, filters)
		require.Equal(t, e, err)
	},
	)
}
