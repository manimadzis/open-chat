package entities

import "time"

type ServerProfile struct {
	UserId
	ServerId
	Nickname string
	JoinTime time.Time
	Roles    []Role
}

type ServerProfileView struct {
	UserId   UserId    `json:"user-id"`
	ServerId ServerId  `json:"server-id"`
	Nickname string    `json:"nickname"`
	JoinTime time.Time `json:"join-time"`
	Roles    []Role    `json:"roles"`
}
