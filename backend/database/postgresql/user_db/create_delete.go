package user_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *UserDatabase) Create(user database.User) error {
	ok, password := user.IsPasswordUpdated()
	if !ok {
		return fmt.Errorf("password not updated")
	}

	if user.Permissions == nil {
		user.Permissions = []string{}
	}

	passwordHash, passwordSalt := hashNewPassword(password)

	query := `
		INSERT INTO public.users (login, name, permissions, password_hash, password_salt)
		VALUES ($1, $2, $3, $4, $5);
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query,
		user.Login,
		user.Name,
		user.Permissions,
		passwordHash,
		passwordSalt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert to db: %w", err)
	}

	return nil
}

func (db *UserDatabase) Delete(login string) error {
	query := `
		DELETE FROM public.users WHERE login = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	res, err := db.pool.Exec(ctx, query, login)
	if err != nil {
		return fmt.Errorf("failed to delete from db: %w", err)
	}

	if res.RowsAffected() == 0 {
		return database.ErrUserNotFound
	}

	return nil
}
