package entities

import (
	"time"
)

type Avatar struct {
	Path string
	Data []byte
}

type Permission uint64

type Role struct {
	Permission Permission
	Server     *Server
}

type Server struct {
	Id     uint64
	Name   string
	Avatar *Avatar
	Users  []User
	Owner  *User
}

type Channel struct {
	Id     uint64
	Name   string
	Server *Server
}

type Sticker struct {
	Id          uint64
	Data        []byte
	StickerPack *StickerPack
}

type File struct {
	Id     uint64
	Name   string
	Path   string
	User   *User
	Server *Server
}

type StickerPack struct {
	Id       uint64
	Name     string
	Stickers []Sticker
}

type Message struct {
	Id      uint64
	Text    string
	Time    time.Time
	Sticker *Sticker
	User    *User
	Channel *Channel
}

type User struct {
	Id               uint64
	Login            string
	Password         string
	Nickname         string
	RegistrationTime time.Time
	Avatar           *Avatar
	Servers          []Server
}

type ServerProfile struct {
	Id       uint64
	Nickname string
	JoinTime time.Time
	User     *User
	Avatar   *Avatar
	Server   *Server
	Roles    []Role
}

type Session struct {
	Id        uint64
	Token     string
	ExpiredAt time.Time
	User      *User
}
