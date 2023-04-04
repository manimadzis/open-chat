package server_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"open-chat/internal/services/role_system"
	"time"
)

type serverService struct {
	serverRepo           services.ServerRepository
	serverProfileChecker services.ServerProfileChecker
}

func (s *serverService) Create(
	ctx context.Context,
	server entities.Server,
) (entities.ServerId, error) {
	var err error
	server.CreationTime = time.Now()

	server.Id, err = s.serverRepo.Create(ctx, server)
	if err != nil {
		return 0, err
	}

	err = s.serverRepo.Join(ctx, server.Id, server.OwnerId)

	return server.Id, err
}

func (s *serverService) Delete(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) error {
	if err := s.serverProfileChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_DELETE_SERVER,
	); err != nil {
		return err
	}
	return s.serverRepo.Delete(ctx, serverId)
}

func (s *serverService) Join(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) error {
	if err := s.serverProfileChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_INVITE_USER,
	); err != nil {
		return err
	}
	return s.serverRepo.Join(ctx, serverId, userId)
}

func (s *serverService) Kick(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) error {
	if err := s.serverProfileChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_KICK_MEMBER,
	); err != nil {
		return err
	}
	return s.serverRepo.Kick(ctx, serverId, userId)
}

func NewServerService(
	serverRepository services.ServerRepository,
	serverProfileChecker services.ServerProfileChecker,
) services.ServerService {
	return &serverService{
		serverRepo:           serverRepository,
		serverProfileChecker: serverProfileChecker,
	}
}
