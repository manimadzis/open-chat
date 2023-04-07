package services

import (
	"errors"
	"fmt"
)

func NewUnknownError(err error) error {
	return &UnknownError{s: fmt.Sprintf("unknown error: %v", err.Error())}
}

type UnknownError struct {
	s string
}

func (u UnknownError) Error() string {
	return u.s
}

var (
	ErrNoSuchPermission           = errors.New("no such permission value")
	ErrNoSuchMessage              = errors.New("no such message")
	ErrNoSuchChannel              = errors.New("no such channel")
	ErrNoSuchToken                = errors.New("no such token")
	ErrUserAlreadyHasToken        = errors.New("user already has token")
	ErrTokenAlreadyExists         = errors.New("token already exists")
	ErrUserDoesntHaveSession      = errors.New("user doesn't have session")
	ErrNoSuchLogin                = errors.New("no such login")
	ErrNoSuchServerProfile        = errors.New("no such server profile")
	ErrLoginAlreadyExists         = errors.New("login already exists")
	ErrInvalidCredentials         = errors.New("invalid credentials")
	ErrNotEnoughPermissions       = errors.New("not enough permissions")
	ErrNoSuchUser                 = errors.New("no such user")
	ErrNoSuchServer               = errors.New("no such server")
	ErrNoSuchStickerPack          = errors.New("no such sticker pack")
	ErrNoSuchSticker              = errors.New("no such sticker pack")
	ErrStickerAlreadyExists       = errors.New("no such sticker pack")
	ErrServerProfileAlreadyExists = errors.New("user already joined server")
)
