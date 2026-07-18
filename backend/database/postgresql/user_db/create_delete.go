package user_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *UserDatabase) CreateUser(user database.User, password string) error {

	passwordHash, passwordSalt := hashPassword(password)

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
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (db *UserDatabase) DeleteUser(login string) error {
	query := `
		DELETE FROM public.users WHERE login = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query, login)

	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
