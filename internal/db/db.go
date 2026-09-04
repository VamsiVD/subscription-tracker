package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, databaseUrl string) (*pgxpool.pool, error) {
	config, err := pgxpool.ParseConfig(databaseUrl)

}
