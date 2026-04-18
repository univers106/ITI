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
	PermSuperUser         = "SuperUser" //не стоит использовать такое в проде, сделал для удобной разработки
)

type Database interface {
	GetUserByID(int) (*User, error)
	GetUserByLogin(string) (*User, error)
	UserAuthentication(login string, password string) (*User, error)
	AddUser(login string, name string, password string) error
	DeleteUser(id int) error
	ChangeUserPassword(user User, password string) error
	UserAddPermissions(user_id int, permission string) error
	UserRemovePermissions(user_id int, permission string) error
	UserCheckPermission(user_id int, permission string) (bool, error)
}

type Post struct {
	ID    int
	Title string
	Body  string
	Date  time.Time
}

type User struct {
	ID          int      "json:\"id\""
	Name        string   "json:\"name\""
	Login       string   "json:\"login\""
	Permissions []string "json:\"permissions\""
}

func (u *User) HasPermission(permission string) bool {
	return slices.Contains(u.Permissions, permission)
}
