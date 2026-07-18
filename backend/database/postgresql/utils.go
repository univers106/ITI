package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func IsTableExists(pool *pgxpool.Pool, tableSchema, tableName string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1
		FROM information_schema.tables
		WHERE table_schema = $1
		AND table_name = $2
	)`
	var exists bool

	reqCtx, cancel := context.WithTimeout(context.Background(), ReqTimeout)
	defer cancel()
	err := pool.QueryRow(reqCtx, query, tableSchema, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
