package entities

import (
	"github.com/google/uuid"
	"time"
)

func NewRole(permissions ...Permission) Role {
	r := Role{}
	for _, perm := range permissions {
		r.Permission |= perm
	}
	return r
}

func NewSession(user *User) *Session {
	return &Session{
		Id:    0,
		Token: uuid.New().String(),
		ExpiredAt: time.Date(5999,
			12, 12, 23,
			59, 59, 0, &time.Location{}),
		User: user,
	}
}
