package user_db

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *UserDatabase) UserAuthentication(login string, password string) (*database.User, error) {
	req := `
		SELECT
			login,
			name,
			permissions,
			password_hash,
			password_salt
		FROM public.users
			WHERE login = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	row := db.pool.QueryRow(ctx, req, login)

	var user User

	err := row.Scan(
		&user.Login,
		&user.Name,
		&user.Permissions,
		&user.PasswordHash,
		&user.PasswordSalt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrIncorrectPassword
		}

		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	if subtle.ConstantTimeCompare(
		hashPassword(password, user.PasswordSalt),
		user.PasswordHash,
	) != 1 {
		return nil, database.ErrIncorrectPassword
	}

	return &user.User, nil
}
