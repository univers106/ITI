// база данных сделана интерфейсом
// в будующем будет реализована более правильная база данных, не на json
// сейчас при аварийном заверешении программы, есть риск смерти всех данных

package database

import (
	"errors"
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

type UserDatabase interface {
	GetByLogin(login string) (*User, error)

	UserAuthentication(login string, password string) (*User, error)

	CreateUser(user User, password string) error
	DeleteUser(login string) error

	ChangeUserPassword(login string, newPassword string) error
	ChangeUserLogin(login string, newLogin string) error
	ChangeUserName(login string, newName string) error

	UserAddPermissions(login string, permission string) error
	UserRemovePermissions(login string, permission string) error
}
