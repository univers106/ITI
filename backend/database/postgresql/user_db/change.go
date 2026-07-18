package user_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
)

func (db *UserDatabase) ChangeUserPassword(login string, newPassword string) error {
	newHash, newSalt := hashNewPassword(newPassword)

	req := "UPDATE users SET password_hash = $1, password_salt = $2 WHERE login = $3"

	affected, err := db.pool.Exec(context.Background(), req, newHash, newSalt, login)
	if err != nil {
		return fmt.Errorf("failed to change user password: %w", err)
	}

	if affected.RowsAffected() == 0 {
		return database.ErrUserNotFound
	}

	return nil
}

func (db *UserDatabase) ChangeUserLogin(login string, newLogin string) error {
	req := "UPDATE users SET login = $2 WHERE login = $1"

	affected, err := db.pool.Exec(context.Background(), req, login, newLogin)
	if err != nil {
		return fmt.Errorf("failed to change user login: %w", err)
	}

	if affected.RowsAffected() == 0 {
		return database.ErrUserNotFound
	}

	return nil
}

func (db *UserDatabase) ChangeUserName(login string, newName string) error {
	req := "UPDATE users SET name = $2 WHERE login = $1"

	affected, err := db.pool.Exec(context.Background(), req, login, newName)
	if err != nil {
		return fmt.Errorf("failed to change user name: %w", err)
	}

	if affected.RowsAffected() == 0 {
		return database.ErrUserNotFound
	}

	return nil
}
