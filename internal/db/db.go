package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func Connect(ctx context.Context, databaseUrl string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", databaseUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}
