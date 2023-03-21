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

type MessageView struct {
	Id      MessageId `json:"id,omitempty"`
	Text    string    `json:"text,omitempty"`
	Time    time.Time `json:"time,omitempty"`
	Sticker *Sticker  `json:"sticker,omitempty"`
	Sender  *User     `json:"sender,omitempty"`
	Channel *Channel  `json:"channel,omitempty"`
}
