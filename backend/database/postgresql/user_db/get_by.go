package user_db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *UserDatabase) GetByLogin(login string) (*database.User, error) {
	query := `
		SELECT login, name, permissions FROM public.users WHERE login = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	row := db.pool.QueryRow(ctx, query, login)

	var user database.User

	err := row.Scan(&user.Login, &user.Name, &user.Permissions)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrUserNotFound
		}

		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	return &user, nil
}
