package storage

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// Init pool and apply migrations
func MustLoadDatabase(databaseURL string, migrationsDir string) *pgxpool.Pool {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnIdleTime = 5 * time.Minute

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := dbpool.Ping(pingCtx); err != nil {
		log.Fatalf("Unable to connect to database (ping failed): %v", err)
	}

	db := stdlib.OpenDBFromPool(dbpool)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	log.Printf("Running migrations from: %s", migrationsDir)

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("Goose migrations failed: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close migration bridge: %v", err)
	}

	log.Println("Database is ready connected and migrated.")
	return dbpool
}
