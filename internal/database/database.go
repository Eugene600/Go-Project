package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(connString string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database %w", err)
	}

	return db, nil
}
