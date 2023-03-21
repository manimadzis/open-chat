package channel_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories"
	"open-chat/internal/services/role_checker"
	"open-chat/internal/services/role_system"
)

type channelService struct {
	channelRepo repositories.ChannelRepository
	roleChecker role_checker.RoleChecker
	serverRepo  repositories.ServerRepository
}

func (s *channelService) Create(ctx context.Context, channel *entities.Channel, userId entities.UserId) error {
	server, err := s.serverRepo.FindByChannelId(ctx, channel.Id)
	if err != nil {
		return err
	}
	if err := s.roleChecker.Check(ctx, userId, server.Id, role_system.PERM_CREATE_CHANNEL); err != nil {
		return err
	}
	return s.channelRepo.Create(ctx, channel)
}

func (s *channelService) Delete(ctx context.Context, channelId entities.ChannelId, userId entities.UserId) error {
	server, err := s.serverRepo.FindByChannelId(ctx, channelId)
	if err != nil {
		return err
	}
	if err := s.roleChecker.Check(ctx, userId, server.Id, role_system.PERM_DELETE_CHANNEL); err != nil {
		return err
	}
	return s.channelRepo.Delete(ctx, channelId)
}

func (s *channelService) FindByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Channel, error) {
	return s.channelRepo.FindByServerId(ctx, serverId)
}

func NewChannelService(channelRepo repositories.ChannelRepository, roleChecker role_checker.RoleChecker, serverRepo repositories.ServerRepository) ChannelService {
	return &channelService{
		channelRepo: channelRepo,
		roleChecker: roleChecker,
		serverRepo:  serverRepo,
	}
}
