package main

import (
	"github.com/WhilsoM/test-go-senior/core/logger"
	"github.com/WhilsoM/test-go-senior/core/storage"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()

	log := logger.NewLogger(cfg.Env)
	defer log.Sync()

	pool := storage.MustLoadDatabase(cfg.DatabaseURL, "./migrations")
	defer pool.Close()

}
