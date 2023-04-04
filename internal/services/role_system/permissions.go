package role_system

import "open-chat/internal/entities"

const (
	PERM_READ_MESSAGE         entities.PermissionValue = 0o00000
	PERM_SEND_MESSAGE         entities.PermissionValue = 0o00002
	PERM_DELETE_MESSAGE       entities.PermissionValue = 0o00004
	PERM_ADD_STICKER          entities.PermissionValue = 0o00010
	PERM_ADD_FILE             entities.PermissionValue = 0o00020
	PERM_CREATE_CHANNEL       entities.PermissionValue = 0o00040
	PERM_DELETE_CHANNEL       entities.PermissionValue = 0o00100
	PERM_KICK_MEMBER          entities.PermissionValue = 0o00200
	PERM_INVITE_USER          entities.PermissionValue = 0o00400
	PERM_CREATE_ROLE          entities.PermissionValue = 0o01000
	PERM_DELETE_ROLE          entities.PermissionValue = 0o02000
	PERM_CHANGE_ROLE          entities.PermissionValue = 0o04000
	PERM_DELETE_SERVER        entities.PermissionValue = 0o10000
	PERM_READ_MESSAGE_HISTORY entities.PermissionValue = 0o20000
)
