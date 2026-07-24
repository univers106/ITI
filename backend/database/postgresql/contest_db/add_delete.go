package contest_db

import (
	"context"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *ContestDatabase) Add(contest database.Contest) error {
	query := `
		INSERT INTO public.contests (name, competitions, is_active)
		VALUES ($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query,
		contest.Name,
		contest.Competitions,
		contest.IsActive,
	)
	if err != nil {
		return database.ErrUnexpected
	}

	return nil
}

func (db *ContestDatabase) Delete(contestName string) error {
	query := `
		DELETE FROM public.contests
		WHERE name = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	_, err := db.pool.Exec(ctx, query, contestName)
	if err != nil {
		return database.ErrUnexpected
	}

	return nil
}
