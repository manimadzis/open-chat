package server_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"open-chat/internal/services/role_system"
	"time"
)

type serverService struct {
	serverRepo        services.ServerRepository
	serverProfileRepo services.ServerProfileRepository
	permissionChecker services.PermissionChecker
	userRepo          services.UserRepository
}

func (s *serverService) Create(
	ctx context.Context,
	server entities.Server,
) (entities.ServerId, error) {
	var err error
	server.CreationTime = time.Now()

	user, err := s.userRepo.FindById(ctx, server.OwnerId)
	if err != nil {
		return 0, err
	}

	server.Id, err = s.serverRepo.Create(ctx, server)
	if err != nil {
		return 0, err
	}

	serverProfile := entities.ServerProfile{
		UserId:   server.OwnerId,
		ServerId: server.Id,
		Nickname: user.Nickname,
		JoinTime: time.Now(),
	}
	if _, err = s.serverProfileRepo.Create(ctx, serverProfile); err != nil {
		return 0, err
	}

	return server.Id, err
}

func (s *serverService) Delete(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) error {
	if err := s.permissionChecker.Check(
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
	if err := s.permissionChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_INVITE_USER,
	); err != nil {
		return err
	}
	user, err := s.userRepo.FindById(ctx, userId)
	if err != nil {
		return err
	}
	serverProfile := entities.ServerProfile{
		UserId:   userId,
		ServerId: serverId,
		Nickname: user.Nickname,
		JoinTime: time.Now(),
	}
	_, err = s.serverProfileRepo.Create(ctx, serverProfile)
	return err
}

func (s *serverService) Kick(
	ctx context.Context,
	serverId entities.ServerId,
	userId entities.UserId,
) error {
	if err := s.permissionChecker.Check(
		ctx,
		userId,
		serverId,
		role_system.PERM_KICK_MEMBER,
	); err != nil {
		return err
	}
	return s.serverProfileRepo.Delete(ctx, entities.ServerProfileId{
		UserId:   userId,
		ServerId: serverId,
	},
	)
}

func NewServerService(
	serverRepo services.ServerRepository,
	serverProfileRepo services.ServerProfileRepository,
	userRepo services.UserRepository,
	permissionChecker services.PermissionChecker,
) services.ServerService {
	return &serverService{
		serverRepo:        serverRepo,
		permissionChecker: permissionChecker,
		serverProfileRepo: serverProfileRepo,
		userRepo:          userRepo,
	}
}
