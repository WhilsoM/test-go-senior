package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RateRepository interface {
	SaveRate(ctx context.Context, ask, bid float64) error
}

type rateRepository struct {
	pool *pgxpool.Pool
}

func NewRateRepository(pool *pgxpool.Pool) RateRepository {
	return &rateRepository{
		pool: pool,
	}
}

func (r *rateRepository) SaveRate(ctx context.Context, ask, bid float64) error {
	query := `INSERT INTO rates (ask, bid) VALUES ($1, $2)`

	_, err := r.pool.Exec(ctx, query, ask, bid)
	if err != nil {
		return fmt.Errorf("failed to insert rate: %w", err)
	}

	return nil
}
