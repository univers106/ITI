package database

import "errors"

var (
	ErrUnexpected         = errors.New("unexpected error")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrPasswordEmpty      = errors.New("password cannot be empty")
	ErrLoginEmpty         = errors.New("login cannot be empty")
	ErrNameEmpty          = errors.New("name cannot be empty")
	ErrIncorrectPassword  = errors.New("incorrect password")
	ErrPermissionNotFound = errors.New("permission not found")
	ErrAlreadyExists      = errors.New("already exists")
)
