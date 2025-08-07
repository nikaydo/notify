package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pg *pgxpool.Pool
}

func InitBD() (Database, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://event:password@localhost:5432/eventbd?sslmode=disable")
	if err != nil {
		return Database{}, err
	}
	return Database{Pg: pool}, nil
}
