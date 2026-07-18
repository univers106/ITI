package user_db

import "github.com/univers106/ITI/database"

type User struct {
	database.User

	ID           string
	PasswordHash []byte
	PasswordSalt []byte
}
