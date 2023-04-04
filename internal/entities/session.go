package entities

import (
	"github.com/google/uuid"
	"time"
)

type SessionToken uuid.UUID

type Session struct {
	Token     SessionToken
	ExpiredAt time.Time
	UserId    UserId
}

func NewInfinitySession(token SessionToken, userId UserId) *Session {
	return &Session{
		Token: token,
		ExpiredAt: time.Date(5999,
			12, 12, 23,
			59, 59, 0, &time.Location{},
		),
		UserId: userId,
	}
}
