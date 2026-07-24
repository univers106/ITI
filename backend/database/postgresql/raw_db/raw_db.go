package raw_db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/univers106/ITI/database/postgresql"
)

var ErrDoesNotExist = fmt.Errorf("raw_data schema does not exist")

type RawDatabase struct {
	pool *pgxpool.Pool
}

func NewRawDatabase(pool *pgxpool.Pool) *RawDatabase {
	tableExists, err := postgresql.IsSchemaExists(pool, "raw_data")
	if err != nil {
		panic(err)
	}

	if !tableExists {
		slog.Warn("raw_data schema does not exist, creating...")

		err := createRawDataSchema(pool)
		if err != nil {
			slog.Error("failed to create raw_data schema", "error", err)
		}
	}

	return &RawDatabase{pool: pool}
}

func CreateRawDataTable(pool *pgxpool.Pool, contest string) error {
	query := fmt.Sprintf(`
	CREATE TABLE raw_data.%s (
		id BIGSERIAL PRIMARY KEY,
		competition VARCHAR(255) NOT NULL,
		student BIGINT NOT NULL,
		score BIGINT NOT NULL,

		CONSTRAINT unique_student_competition_%s UNIQUE (competition, student)
	);`, contest, contest)

	reqCtx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := pool.Exec(reqCtx, query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу: %w", err)
	}

	return nil
}

func createRawDataSchema(pool *pgxpool.Pool) error {
	query := `CREATE SCHEMA raw_data;`

	reqCtx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := pool.Exec(reqCtx, query)
	if err != nil {
		return fmt.Errorf("не удалось создать схему: %w", err)
	}

	return nil
}
