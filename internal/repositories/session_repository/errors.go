package session_repository

import "errors"

var (
	ErrNoSuchToken           = errors.New("no such token")
	ErrUserAlreadyHasToken   = errors.New("user already has token")
	ErrTokenAlreadyExists    = errors.New("token already exists")
	ErrUserDoesntHaveSession = errors.New("user doesn't have session")
)
