package entities

import "time"

type Server struct {
	Id           ServerId
	Name         string
	OwnerId      UserId
	CreationTime time.Time
}

type ServerView struct {
	Id    ServerId `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Users []User   `json:"users,omitempty"`
	Owner *User    `json:"owner,omitempty"`
}
