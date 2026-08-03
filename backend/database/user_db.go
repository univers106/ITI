package database

const (
	PermUsersManipulation = "UsersManipulation"
	PermSuperUser         = "SuperUser" // не стоит использовать такое в проде, сделал для удобной разработки
)

type UserDatabase interface {
	GetByLogin(login string) (*User, error)
	GetAll() ([]User, error)

	UserAuthentication(login string, password string) (*User, error)

	Create(user User) error
	Delete(login string) error

	Change(login string, user User) error
}
