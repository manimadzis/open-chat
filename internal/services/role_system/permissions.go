package role_system

const (
	PERM_READ_MESSAGE   = 0o0000
	PERM_SEND_MESSAGE   = 0o0002
	PERM_DELETE_MESSAGE = 0o0004

	PERM_ADD_STICKER = 0o0010
	PERM_ADD_FILE    = 0o0020

	PERM_CREATE_CHANNEL = 0o0040
	PERM_DELETE_CHANNEL = 0o0100

	PERM_KICK_MEMBER = 0o0200
	PERM_INVITE_USER = 0o0400

	PERM_CREATE_ROLE = 0o0100
	PERM_DELETE_ROLE = 0o0100
	PERM_CHANGE_ROLE = 0o0200

	PERM_DELETE_SERVER = 0o0400

	PERM_READ_MESSAGE_HISTORY = 0o1000
)