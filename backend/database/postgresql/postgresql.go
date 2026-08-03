package postgresql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	ReqTimeout = 3 * time.Second
)

func NewPool(databaseURL string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), ReqTimeout)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		panic(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}

	return pool
}
