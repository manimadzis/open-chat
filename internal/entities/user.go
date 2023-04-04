package entities

import "time"

type User struct {
	Id               UserId
	Login            string
	Password         string
	Nickname         string
	RegistrationTime time.Time
}
