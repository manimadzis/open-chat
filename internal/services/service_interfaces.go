package services

import (
	"context"
	"open-chat/internal/entities"
)

type ChannelService interface {
	// Create returns ErrNotEnoughPermissions if user doesn't have enough permissions.
	Create(
		ctx context.Context,
		channel entities.Channel,
	) (
		entities.ChannelId,
		error,
	)

	// Delete returns ErrNoSuchChannelId if given channel doesn't exists.
	// If user doesn't have enough permissions it returns ErrNotEnoughPermissions.
	Delete(
		ctx context.Context,
		channelId entities.ChannelId,
		userId entities.UserId,
	) error

	// FindByServerId returns ErrNoSuchServerProfile if given user doesn't have profile on the server.
	// If user doesn't have enough permissions it returns ErrNotEnoughPermissions.
	FindByServerId(
		ctx context.Context,
		serverId entities.ServerId,
		userId entities.UserId,
	) (
		[]entities.Channel,
		error,
	)
}

type MessageService interface {
	// Send sends message. If user doesn't have sent permission it returns ErrNotEnoughPermissions
	Send(
		ctx context.Context,
		message entities.Message,
	) (
		entities.MessageId,
		error,
	)
	// Delete deletes message. If user doesnt have
	Delete(
		ctx context.Context,
		messageId entities.MessageId,
		userId entities.UserId,
	) error
	FindInChat(
		ctx context.Context,
		userId entities.UserId,
		channelId entities.ChannelId,
		filters entities.MessageFiltersDTO,
	) (
		[]entities.Message,
		error,
	)
}

type RoleService interface {
	Create(
		ctx context.Context,
		role entities.Role,
	) (
		entities.RoleId,
		error,
	)
	Delete(
		ctx context.Context,
		roleId entities.RoleId,
		userId entities.UserId,
		serverId entities.ServerId,
	) error
	Change(
		ctx context.Context,
		role entities.Role,
		userId entities.UserId,
		serverId entities.ServerId,
	) error
	FindByServer(
		ctx context.Context,
		serverId entities.ServerId,
		userId entities.UserId,
	) (
		[]entities.Role,
		error,
	)
}

type PermissionChecker interface {
	// Check returns ErrNoSuchServerProfile if user doesn't have server profile.
	// If user doesn't have enough permissions it returns ErrNotEnoughPermissions.
	// Could be used without permissions to check user has been joined the server
	// You should use it in all services dealing with server otherwise the user can get
	// information from server without joining it
	Check(
		ctx context.Context,
		userId entities.UserId,
		serverId entities.ServerId,
		permissions ...entities.PermissionValue,
	) error
}

type ServerService interface {
	Create(
		ctx context.Context,
		server entities.Server,
	) (
		entities.ServerId,
		error,
	)
	Delete(
		ctx context.Context,
		serverId entities.ServerId,
		userId entities.UserId,
	) error
	Join(
		ctx context.Context,
		serverId entities.ServerId,
		userId entities.UserId,
	) error
	Kick(
		ctx context.Context,
		serverId entities.ServerId,
		userId entities.UserId,
	) error
}

type AuthService interface {
	SignUp(
		ctx context.Context,
		user entities.User,
	) (
		*entities.Session,
		error,
	)
	SignIn(
		ctx context.Context,
		user entities.User,
	) (
		*entities.Session,
		error,
	)
	LogOut(
		ctx context.Context,
		sessionToken entities.SessionToken,
	) error
	FindSessionByToken(
		ctx context.Context,
		sessionToken entities.SessionToken,
	) (
		*entities.Session,
		error,
	)
}

type StickerService interface {
	CreateStickerPack(
		ctx context.Context,
		stickerPack entities.StickerPack,
	) (
		entities.StickerPackId,
		error,
	)
	DeleteStickerPack(
		ctx context.Context,
		stickerPackId entities.StickerPackId,
	) error
	FindStickersByStickerPackId(
		ctx context.Context,
		stickerPackId entities.StickerPackId,
	) (
		[]entities.Sticker,
		error,
	)
}
