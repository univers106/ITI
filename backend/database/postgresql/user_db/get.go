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

func (db *UserDatabase) GetAll() ([]database.User, error) {
	query := `
		SELECT login, name, permissions FROM public.users;
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []database.User
	for rows.Next() {
		var user database.User
		err := rows.Scan(&user.Login, &user.Name, &user.Permissions)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
