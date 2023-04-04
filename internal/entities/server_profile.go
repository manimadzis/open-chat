package entities

import "time"

type ServerProfile struct {
	UserId
	ServerId
	Nickname string
	JoinTime time.Time
	Roles    []Role
}
