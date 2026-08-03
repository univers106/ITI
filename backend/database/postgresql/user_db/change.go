package user_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *UserDatabase) Change(login string, user database.User) error {
	var newHash, newSalt []byte

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	isPasswordUpdated, newPassword := user.IsPasswordUpdated()
	if isPasswordUpdated {
		newHash, newSalt = hashNewPassword(newPassword)
		req := "UPDATE users SET login = $1, permissions = $2, password_hash = $3, password_salt = $4 WHERE login = $5"

		affected, err := db.pool.Exec(ctx, req, user.Login, user.Permissions, newHash, newSalt, login)
		if err != nil {
			return fmt.Errorf("failed to change user password: %w", err)
		}

		if affected.RowsAffected() == 0 {
			return database.ErrUserNotFound
		}
		return nil
	}

	req := "UPDATE users SET login = $1, permissions = $2 WHERE login = $3"
	affected, err := db.pool.Exec(ctx, req, user.Login, user.Permissions, login)
	if err != nil {
		return fmt.Errorf("failed to change user: %w", err)
	}

	if affected.RowsAffected() == 0 {
		return database.ErrUserNotFound
	}
	return nil
}
