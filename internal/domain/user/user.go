package user

import "errors"

const DEFAULT_BALANCE = 1000

type User struct {
	ID           string
	Username     string
	PasswordHash string
	Balance      uint
}

var ErrUserNotFound = errors.New("user not found")
