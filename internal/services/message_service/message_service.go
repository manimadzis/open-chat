package message_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/message_repository"
	"open-chat/internal/repositories/user_repository"
	"open-chat/internal/services/role_checker"
	"open-chat/internal/services/role_system"
)

type messageService struct {
	messageRepo message_repository.MessageRepository
	userRepo    user_repository.UserRepository
	roleChecker role_checker.RoleChecker
}

func (m *messageService) Send(ctx context.Context, message *entities.Message) error {
	if err := m.roleChecker.Check(ctx, message.User, message.Channel.Server, role_system.PERM_SEND_MESSAGE); err != nil {
		return err
	}
	return m.messageRepo.Create(ctx, message)
}

func (m *messageService) Delete(ctx context.Context, message *entities.Message) error {
	if err := m.roleChecker.Check(ctx, message.User, message.Channel.Server, role_system.PERM_SEND_MESSAGE); err != nil {
		return err
	}
	return m.messageRepo.Delete(ctx, message)
}

func (m *messageService) Find(ctx context.Context, user *entities.User, channel *entities.Channel, filters *entities.MessageFiltersDTO) ([]*entities.Message, error) {
	if err := m.roleChecker.Check(ctx, user, channel.Server, role_system.PERM_READ_MESSAGE); err != nil {
		return nil, err
	}

	messages, err := m.messageRepo.FindByChannel(ctx, channel, filters.Offset, filters.Count)
	if err != nil {
		return nil, err
	}

	if err := m.roleChecker.Check(ctx, user, channel.Server, role_system.PERM_READ_MESSAGE_HISTORY); err == nil {
		serverProfile, err := m.userRepo.FindServerProfileByIds(ctx, user, channel.Server)
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

func NewMessageService(repository message_repository.MessageRepository, roleChecker role_checker.RoleChecker) MessageService {
	return &messageService{
		messageRepo: repository,
		roleChecker: roleChecker,
	}
}
