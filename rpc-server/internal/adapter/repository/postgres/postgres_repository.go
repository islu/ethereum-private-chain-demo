package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	connPool *pgxpool.Pool
}

func NewPostgresRepository(ctx context.Context, connPool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		connPool: connPool,
	}
}
