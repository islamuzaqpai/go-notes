package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	connStr := "postgres://postgres:kuralai@localhost:5432/postgres"
	con, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}
	return con, err
}
