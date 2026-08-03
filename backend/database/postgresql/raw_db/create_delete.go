package raw_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *RawDatabase) Add(contest string, raw database.Raw) error {
	query := fmt.Sprintf(`
		INSERT INTO raw_data.%s (competition, student, score)
		VALUES ($1, $2, $3);
	`, contest)

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query,
		raw.CompetitionId,
		raw.Student,
		raw.Score,
	)
	if err != nil {
		return database.ErrUnexpected
	}

	return nil
}

func (db *RawDatabase) Delete(contest string, raw database.Raw) error {
	query := fmt.Sprintf(`
		DELETE FROM raw_data.%s
		WHERE
			student = $1 AND
		 	score = $2;
	`, contest)

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query, raw.Student, raw.Score)
	if err != nil {
		return database.ErrUnexpected
	}

	return nil
}
