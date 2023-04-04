package entities

import "time"

type Channel struct {
	Id           ChannelId
	Name         string
	CreationTime time.Time
	ServerId     ServerId
	CreatorId    UserId
}
