package handlers

import (
	"open-chat/internal/entities"
	"time"
)

type MessageView struct {
	Id      entities.MessageId `json:"id,omitempty"`
	Text    string             `json:"text,omitempty"`
	Time    time.Time          `json:"time,omitempty"`
	Sticker *entities.Sticker  `json:"sticker,omitempty"`
	Sender  *entities.User     `json:"sender,omitempty"`
	Channel *entities.Channel  `json:"channel,omitempty"`
}

type ChannelView struct {
	Id        entities.ChannelId `json:"id"`
	Name      string             `json:"name"`
	ServerId  entities.ServerId  `json:"server-id"`
	CreatorId entities.UserId    `json:"creator-id"`
}

type FileView struct {
	Id   entities.FileId `json:"id"`
	Path string          `json:"path"`
}

type RoleView struct {
	Id              entities.RoleId          `json:"id"`
	Name            string                   `json:"name"`
	PermissionValue entities.PermissionValue `json:"permission-value"`
	CreatedAt       time.Time                `json:"create-at"`
	Permissions     []entities.Permission    `json:"permissions"`
	Creator         *entities.User           `json:"creator"`
}

type ServerView struct {
	Id    entities.ServerId `json:"id,omitempty"`
	Name  string            `json:"name,omitempty"`
	Users []entities.User   `json:"users,omitempty"`
	Owner *entities.User    `json:"owner,omitempty"`
}

type ServerProfileView struct {
	UserId   entities.UserId   `json:"user-id"`
	ServerId entities.ServerId `json:"server-id"`
	Nickname string            `json:"nickname"`
	JoinTime time.Time         `json:"join-time"`
	Roles    []entities.Role   `json:"roles"`
}

type SessionView struct {
	Token     entities.SessionToken `json:"token"`
	ExpiredAt time.Time             `json:"expired-at"`
}

type StickerView struct {
	Id  entities.StickerId `json:"id"`
	URL string             `json:"url"`
}

type UserView struct {
	Id       entities.UserId `json:"id"`
	Nickname string          `json:"nickname"`
}
