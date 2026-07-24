package contest_db

import (
	"context"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *ContestDatabase) Edit(contest string, contestData database.Contest) error {
	query := `
		UPDATE public.contests
		SET
			name = $2,
			competitions = $3,
			is_active = $4
		WHERE name = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query, contest,
		contestData.Name,
		contestData.Competitions,
		contestData.IsActive,
	)
	if err != nil {
		return database.ErrUnexpected
	}

	return nil
}
