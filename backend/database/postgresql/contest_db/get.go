package contest_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *ContestDatabase) Get(contestName string) (database.Contest, error) {
	query := `
		SELECT name, competitions, is_active FROM public.contests
		WHERE name = $1
		`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	row := db.pool.QueryRow(ctx, query, contestName)

	var contest database.Contest

	err := row.Scan(&contest.Name, &contest.Competitions, &contest.IsActive)
	if err != nil {
		return contest, fmt.Errorf("failed to get contest: %w", err)
	}

	return contest, nil
}

func (db *ContestDatabase) List() ([]database.Contest, error) {
	query := `
		SELECT name, competitions, is_active FROM public.contests
		`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get contests: %w", err)
	}

	var contests []database.Contest

	for rows.Next() {
		var contest database.Contest

		err := rows.Scan(&contest.Name, &contest.Competitions, &contest.IsActive)
		if err != nil {
			return nil, fmt.Errorf("failed to scan contest: %w", err)
		}

		contests = append(contests, contest)
	}

	return contests, nil
}

func (db *ContestDatabase) ActiveContests() ([]database.Contest, error) {
	query := `
		SELECT name, competitions, is_active FROM public.contests
		WHERE is_active = true
		`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get active contests: %w", err)
	}

	var contests []database.Contest

	for rows.Next() {
		var contest database.Contest

		err := rows.Scan(&contest.Name, &contest.Competitions, &contest.IsActive)
		if err != nil {
			return nil, fmt.Errorf("failed to scan contest: %w", err)
		}

		contests = append(contests, contest)
	}

	return contests, nil
}
