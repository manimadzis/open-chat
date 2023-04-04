package message_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"open-chat/internal/services/role_system"
)

type messageService struct {
	messageRepo          services.MessageRepository
	serverRepo           services.ServerRepository
	serverProfileChecker services.ServerProfileChecker
}

func (m *messageService) Send(
	ctx context.Context,
	message entities.Message,
) (entities.MessageId, error) {
	server, err := m.serverRepo.FindByChannelId(ctx, message.ChannelId)
	if err != nil {
		return 0, err
	}
	if err := m.serverProfileChecker.Check(
		ctx,
		message.SenderId,
		server.Id,
		role_system.PERM_SEND_MESSAGE,
	); err != nil {
		return 0, err
	}
	return m.messageRepo.Create(ctx, message)
}

func (m *messageService) Delete(
	ctx context.Context,
	messageId entities.MessageId,
	userId entities.UserId,
) error {
	server, err := m.serverRepo.FindByMessageId(ctx, messageId)
	if err != nil {
		return err
	}
	if err := m.serverProfileChecker.Check(
		ctx,
		userId,
		server.Id,
		role_system.PERM_DELETE_MESSAGE,
	); err != nil {
		return err
	}
	return m.messageRepo.Delete(ctx, messageId)
}

func (m *messageService) FindInChat(
	ctx context.Context,
	userId entities.UserId,
	channelId entities.ChannelId,
	filters entities.MessageFiltersDTO,
) ([]entities.Message, error) {
	server, err := m.serverRepo.FindByChannelId(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if err := m.serverProfileChecker.Check(
		ctx,
		userId,
		server.Id,
		role_system.PERM_READ_MESSAGE,
	); err != nil {
		return nil, err
	}

	messages, err := m.messageRepo.FindByChannel(
		ctx,
		channelId,
		filters.Offset,
		filters.Count,
	)
	if err != nil {
		return nil, err
	}

	if err := m.serverProfileChecker.Check(
		ctx,
		userId,
		server.Id,
		role_system.PERM_READ_MESSAGE_HISTORY,
	); err != nil {
		serverProfile, err := m.serverRepo.FindServerProfileByIds(
			ctx,
			server.Id,
			userId,
		)
		if err != nil {
			return nil, err
		}

		ii := len(messages)
		for i, message := range messages {
			if message.Time.Before(serverProfile.JoinTime) {
				ii = i
				break
			}
		}
		messages = messages[:ii]
	}

	return messages, nil
}

func NewMessageService(
	messageRepo services.MessageRepository,
	serverRepo services.ServerRepository,
	serverProfileChecker services.ServerProfileChecker,
) services.MessageService {
	return &messageService{
		messageRepo:          messageRepo,
		serverRepo:           serverRepo,
		serverProfileChecker: serverProfileChecker,
	}
}
