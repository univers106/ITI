package user_db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/univers106/ITI/database/postgresql"
)

type UserDatabase struct {
	pool *pgxpool.Pool
}

func NewUserDatabase(pool *pgxpool.Pool) *UserDatabase {
	tableExists, _ := postgresql.IsTableExists(pool, "public", "users")
	if !tableExists {
		slog.Warn("users table does not exist, creating...")

		err := createUsersTable(pool)
		if err != nil {
			slog.Error("failed to create users table", "error", err)
		}
	}

	return &UserDatabase{pool: pool}
}

func createUsersTable(pool *pgxpool.Pool) error {
	query := `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(50) NOT NULL UNIQUE,
		name VARCHAR(100) NOT NULL,
		permissions TEXT[] NOT NULL DEFAULT '{}',
		password_hash TEXT NOT NULL
	);`

	reqCtx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()
	_, err := pool.Exec(reqCtx, query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу: %w", err)
	}

	return nil
}
