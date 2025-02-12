package auth

import "errors"

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrInvalidPassword = errors.New("invalid password")
)
