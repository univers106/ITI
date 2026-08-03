package competition_db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/univers106/ITI/database/postgresql"
)

type CompetitionDatabase struct {
	pool *pgxpool.Pool
}

func NewCompetitionDatabase(pool *pgxpool.Pool) *CompetitionDatabase {
	tableExists, err := postgresql.IsTableExists(pool, "public", "competitions")
	if err != nil {
		panic(err)
	}

	if !tableExists {
		slog.Warn("contests table does not exist, creating...")

		err := createContestsTable(pool)
		if err != nil {
			slog.Error("failed to create competitions table", "error", err)
		}
	}

	return &CompetitionDatabase{pool: pool}
}

func createContestsTable(pool *pgxpool.Pool) error {
	query := `
	CREATE TABLE public.competitions (
		id BIGSERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		competitions INT[] NOT NULL DEFAULT '{}',
		is_active BOOLEAN NOT NULL DEFAULT TRUE,
	);`

	reqCtx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := pool.Exec(reqCtx, query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу: %w", err)
	}

	return nil
}
