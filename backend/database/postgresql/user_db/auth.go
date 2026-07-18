package user_db

import (
	"context"
	"crypto/subtle"
	"errors"

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
		&user.User.Login,
		&user.User.Name,
		&user.User.Permissions,
		&user.PasswordHash,
		&user.PasswordSalt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAuthInvalid
		}
		return nil, err
	}

	if subtle.ConstantTimeCompare(hashPassword(password, user.PasswordSalt), user.PasswordHash) != 1 {
		return nil, ErrAuthInvalid
	}

	return &user.User, nil
}
