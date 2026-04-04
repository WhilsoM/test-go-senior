package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type RateRepository interface {
	SaveRate(ctx context.Context, ask, bid float64) error
}

type rateRepository struct {
	pool *pgxpool.Pool
	log  *zap.Logger
}

func NewRateRepository(pool *pgxpool.Pool, log *zap.Logger) RateRepository {
	return &rateRepository{
		pool: pool,
		log:  log,
	}
}

func (r *rateRepository) SaveRate(ctx context.Context, ask, bid float64) error {
	query := `INSERT INTO rates (ask, bid) VALUES ($1, $2)`

	r.log.Info("SaveRate started")

	_, err := r.pool.Exec(ctx, query, ask, bid)
	if err != nil {
		r.log.Error("SaveRate have an error", zap.Error(err))
		return fmt.Errorf("failed to insert rate: %w", err)
	}

	r.log.Info("SaveRate successfully saved")

	return nil
}
