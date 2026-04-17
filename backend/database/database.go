package database

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserExists        = errors.New("user already exists")
	ErrPasswordEmpty     = errors.New("password cannot be empty")
	ErrLoginEmpty        = errors.New("login cannot be empty")
	ErrNameEmpty         = errors.New("name cannot be empty")
	ErrIncorrectPassword = errors.New("incorrect password")
)

type Database interface {
	GetUserByID(int) (User, error)
	GetUserByLogin(string) (User, error)
	UserAuthentication(login string, password string) (User, error)
	AddUser(login string, name string, password string) error
	DeleteUser(id int) error
	ChangeUserPassword(user User, password string) error
}

type Post struct {
	ID    int
	Title string
	Body  string
	Date  time.Time
}

type User struct {
	ID    int    "json:\"id\""
	Name  string "json:\"name\""
	Login string "json:\"login\""
}
