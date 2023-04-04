package entities

import "time"

type Message struct {
	Id       MessageId
	Text     string
	Time     time.Time
	SenderId UserId
	ChannelId
	Sticker *Sticker
}
