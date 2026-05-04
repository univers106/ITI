// база данных сделана интерфейсом
// в будующем будет реализована более правильная база данных, не на json
// сейчас при аварийном заверешении программы, есть риск смерти всех данных

package database

import (
	"errors"
	"slices"
	"time"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrPasswordEmpty      = errors.New("password cannot be empty")
	ErrLoginEmpty         = errors.New("login cannot be empty")
	ErrNameEmpty          = errors.New("name cannot be empty")
	ErrIncorrectPassword  = errors.New("incorrect password")
	ErrPermissionNotFound = errors.New("permission not found")
	ErrAlreadyExists      = errors.New("already exists")
)

const (
	PermUsersManipulation = "UsersManipulation"
	PermSuperUser         = "SuperUser" // не стоит использовать такое в проде, сделал для удобной разработки
)

type Database interface {
	GetUserByID(id int) (*User, error)
	GetUserByLogin(login string) (*User, error)
	UserAuthentication(login string, password string) (*User, error)
	CreateUser(login string, name string, password string) error
	DeleteUser(id int) error

	ChangeUserPassword(userId int, password string) error
	ChangeUserLogin(userId int, login string) error
	ChangeUserName(userId int, name string) error

	UserAddPermissions(userId int, permission string) error
	UserRemovePermissions(userId int, permission string) error
	UserCheckPermission(userId int, permission string) (bool, error)
}

type Post struct {
	ID    int
	Title string
	Body  string
	Date  time.Time
}

type User struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Login       string   `json:"login"`
	Permissions []string `json:"permissions"`
}

func (u *User) HasPermission(permission string) bool {
	if slices.Contains(u.Permissions, PermSuperUser) {
		return true
	}

	return slices.Contains(u.Permissions, permission)
}
