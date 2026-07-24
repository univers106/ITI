package database

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

	UserAddPermission(login string, permission string) error
	UserRemovePermission(login string, permission string) error
}
