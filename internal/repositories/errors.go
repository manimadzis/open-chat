package repositories

import (
	"errors"
	"fmt"
)

func UnknownError(err error) error {
	return fmt.Errorf("unknown error: %v", err.Error())
}

var (
	ErrNoSuchPermission      = errors.New("no such permission value")
	ErrNoSuchMessage         = errors.New("no such message")
	ErrNoSuchChannel         = errors.New("no such channel")
	ErrNoSuchToken           = errors.New("no such token")
	ErrUserAlreadyHasToken   = errors.New("user already has token")
	ErrTokenAlreadyExists    = errors.New("token already exists")
	ErrUserDoesntHaveSession = errors.New("user doesn't have session")
	ErrNoSuchLogin           = errors.New("no such login")
)
