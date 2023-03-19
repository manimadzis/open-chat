package user_service

import "errors"

var (
	ErrLoginAlreadyUsed = errors.New("login already used")
)
