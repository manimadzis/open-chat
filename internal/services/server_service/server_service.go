package server_service

import (
	"context"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/server_repository"
	"open-chat/internal/services/role_checker"
	"open-chat/internal/services/role_system"
)

type serverService struct {
	serverRepo  server_repository.ServerRepository
	roleChecker role_checker.RoleChecker
}

func (s *serverService) Create(ctx context.Context, server *entities.Server, user *entities.User) error {
	if err := s.serverRepo.Create(ctx, server, user); err != nil {
		return err
	}
	return s.serverRepo.Join(ctx, server, user)
}

func (s *serverService) Delete(ctx context.Context, server *entities.Server, user *entities.User) error {
	if err := s.roleChecker.Check(ctx, user, server, role_system.PERM_DELETE_SERVER); err != nil {
		return err
	}
	return s.serverRepo.Delete(ctx, server, user)
}

func (s *serverService) Join(ctx context.Context, server *entities.Server, user *entities.User) error {
	if err := s.roleChecker.Check(ctx, user, server, role_system.PERM_INVITE_USER); err != nil {
		return err
	}
	return s.serverRepo.Join(ctx, server, user)
}

func (s *serverService) Kick(ctx context.Context, server *entities.Server, user *entities.User) error {
	if err := s.roleChecker.Check(ctx, user, server, role_system.PERM_KICK_MEMBER); err != nil {
		return err
	}
	return s.serverRepo.Join(ctx, server, user)
}

func NewServerService(serverRepository server_repository.ServerRepository, roleChecker role_checker.RoleChecker) ServerService {
	return &serverService{
		serverRepo:  serverRepository,
		roleChecker: roleChecker,
	}
}
