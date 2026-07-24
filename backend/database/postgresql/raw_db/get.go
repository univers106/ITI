package raw_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *RawDatabase) Get(
	contest string,
	competition string,
	student int,
) ([]database.Raw, error) {
	query := fmt.Sprintf(`
		SELECT
			competition, student, score
		FROM raw_data.%s
		WHERE
			competition = $1 AND
			student = $2;
	`, contest)

	return db.getter(query, competition, student)
}

func (db *RawDatabase) GetByCompetition(
	contest string,
	competition string,
) ([]database.Raw, error) {
	query := fmt.Sprintf(`
		SELECT
			competition, student, score
		FROM raw_data.%s
		WHERE
			competition = $1;
	`, contest)

	return db.getter(query, competition)
}

func (db *RawDatabase) GetByStudent(contest string, student int) ([]database.Raw, error) {
	query := fmt.Sprintf(`
		SELECT
			competition, student, score
		FROM raw_data.%s
		WHERE
			student = $1;
	`, contest)

	return db.getter(query, student)
}

func (db *RawDatabase) getter(query string, args ...any) ([]database.Raw, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	rows, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	defer rows.Close()

	result := []database.Raw{}

	for rows.Next() {
		var raw database.Raw

		err := rows.Scan(&raw.CompetitionId, &raw.Student, &raw.Score)
		if err != nil {
			return nil, fmt.Errorf("failed to scan raw: %w", err)
		}

		result = append(result, raw)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return result, nil
}
