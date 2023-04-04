package entities

import "time"

type Server struct {
	Id           ServerId
	Name         string
	OwnerId      UserId
	CreationTime time.Time
}
