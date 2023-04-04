package services

import (
	"context"
	"open-chat/internal/entities"
)

type ChannelRepository interface {
	// Create creates channel and returns its id.
	Create(ctx context.Context, channel entities.Channel) (entities.ChannelId, error)

	// Delete deletes channel with given id. If channel doesn't exist it returns ErrNoSuchChannelId.
	Delete(ctx context.Context, channelId entities.ChannelId) error

	// FindByServerId returns slice of channels for given server
	FindByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Channel, error)
}

type MessageRepository interface {
	// Create creates message. Set Id field in given struct.
	Create(ctx context.Context, message entities.Message) (entities.MessageId, error)

	// Delete deletes message with given id. It returns ErrNoSuchMessage if given id doesn't exist.
	Delete(ctx context.Context, messageId entities.MessageId) error

	// FindByChannel returns slice of messages for given channel. If offset is too big it returns empty slice without errors.
	FindByChannel(ctx context.Context, channelId entities.ChannelId, offset, count uint64) ([]entities.Message, error)
}

type RoleRepository interface {
	Create(ctx context.Context, role entities.Role) (entities.RoleId, error)

	Delete(ctx context.Context, roleId entities.RoleId) error

	Change(ctx context.Context, role *entities.Role) error

	// FindRolesByServerId returns role by server id. If no server with given id it returns ErrNoSuchServer
	FindRolesByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Role, error)

	// FindPermissionsByValue returns slice of permissions by its value. If PermissionValue is invalid it returns ErrNoSuchPermission.
	FindPermissionsByValue(ctx context.Context, permission []entities.PermissionValue) ([]entities.Permission, error)
}

type ServerRepository interface {
	Create(ctx context.Context, server entities.Server) (entities.ServerId, error)

	Delete(ctx context.Context, serverId entities.ServerId) error

	Join(ctx context.Context, serverId entities.ServerId, userId entities.UserId) error
	Kick(ctx context.Context, serverId entities.ServerId, userId entities.UserId) error

	FindByMessageId(ctx context.Context, messageId entities.MessageId) (*entities.Server, error)
	// FindByChannelId returns ErrNoSuchChannel if channel doesn't exists
	FindByChannelId(ctx context.Context, channelId entities.ChannelId) (*entities.Server, error)

	// FindServerProfileByIds returns ErrNoSuchServerProfile if given user doesn't have profile on the server (user didn't join).
	FindServerProfileByIds(ctx context.Context,
		serverId entities.ServerId,
		userId entities.UserId,
	) (*entities.ServerProfile, error)
}

type UserRepository interface {
	// Create returns ErrLoginAlreadyExists if login already exists.
	Create(ctx context.Context, user entities.User) (entities.UserId, error)

	FindById(ctx context.Context, userId entities.UserId) (*entities.User, error)

	// FindByLogin returns ErrNoSuchLogin if given login doesn't exist.
	FindByLogin(ctx context.Context, login string) (*entities.User, error)
}

type SessionRepository interface {
	// Create write session. If token already has used it returns ErrTokenAlreadyExists
	Create(ctx context.Context, session entities.Session) error

	// DeleteByToken delete session by token. If no session found by token it returns ErrNoSuchToken
	DeleteByToken(ctx context.Context, session entities.SessionToken) error

	// FindByToken find session by token. If no session found by token it returns ErrNoSuchToken
	FindByToken(ctx context.Context, session entities.SessionToken) (*entities.Session, error)

	// FindByUserId find session by user id. If no session found it returns ErrUserDoesntHaveSession
	FindByUserId(ctx context.Context, session *entities.Session) (*entities.Session, error)
}
