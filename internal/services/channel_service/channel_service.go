package channel_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"open-chat/internal/services/role_system"
)

type channelService struct {
	channelRepo       services.ChannelRepository
	permissionChecker services.PermissionChecker
	serverRepo        services.ServerRepository
}

func (s channelService) Create(
	ctx context.Context,
	channel entities.Channel,
) (entities.ChannelId, error) {
	if err := s.permissionChecker.Check(
		ctx,
		channel.CreatorId,
		channel.ServerId,
		role_system.PERM_CREATE_CHANNEL,
	); err != nil {
		return 0, err
	}
	return s.channelRepo.Create(ctx, channel)
}

func (s channelService) Delete(
	ctx context.Context,
	channelId entities.ChannelId,
	userId entities.UserId,
) error {
	server, err := s.serverRepo.FindByChannelId(ctx, channelId)
	if err != nil {
		return err
	}
	if err := s.permissionChecker.Check(
		ctx,
		userId,
		server.Id,
		role_system.PERM_DELETE_CHANNEL,
	); err != nil {
		return err
	}
	return s.channelRepo.Delete(ctx, channelId)
}

func (s channelService) FindByServerId(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) ([]entities.Channel, error) {
	if err := s.permissionChecker.Check(ctx, userId, serverId); err != nil {
		return nil, err
	}
	return s.channelRepo.FindByServerId(ctx, serverId)
}

func NewChannelService(
	channelRepo services.ChannelRepository,
	permissionChecker services.PermissionChecker,
	serverRepo services.ServerRepository,
) services.ChannelService {
	return &channelService{
		channelRepo:       channelRepo,
		permissionChecker: permissionChecker,
		serverRepo:        serverRepo,
	}
}
