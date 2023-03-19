package channel_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/channel_repository"
	"open-chat/internal/services/role_checker"
	"open-chat/internal/services/role_system"
)

type channelService struct {
	channelRepo channel_repository.ChannelRepository
	roleChecker role_checker.RoleChecker
}

func (s *channelService) Create(ctx context.Context, channel *entities.Channel, user *entities.User) error {
	if err := s.roleChecker.Check(ctx, user, channel.Server, role_system.PERM_CREATE_CHANNEL); err != nil {
		return err
	}
	return s.channelRepo.Create(ctx, channel, user)
}

func (s *channelService) Delete(ctx context.Context, channel *entities.Channel, user *entities.User) error {
	if err := s.roleChecker.Check(ctx, user, channel.Server, role_system.PERM_DELETE_CHANNEL); err != nil {
		return err
	}
	return s.channelRepo.Delete(ctx, channel, user)
}

func (s *channelService) FindByServer(ctx context.Context, server *entities.Server, user *entities.User) ([]*entities.Channel, error) {
	return s.channelRepo.FindByServer(ctx, server, user)
}

func NewChannelService(channelRepo channel_repository.ChannelRepository, roleChecker role_checker.RoleChecker) ChannelService {
	return &channelService{
		channelRepo: channelRepo,
		roleChecker: roleChecker,
	}
}
