package entities

import "time"

type Channel struct {
	Id           ChannelId
	Name         string
	CreationTime time.Time
	ServerId     ServerId
	CreatorId    UserId
}

type ChannelView struct {
	Id        ChannelId `json:"id"`
	Name      string    `json:"name"`
	ServerId  ServerId  `json:"server-id"`
	CreatorId UserId    `json:"creator-id"`
}
