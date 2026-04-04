package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DatabaseURL  string `env:"DATABASE_URL" env-default:"postgres://user:pass@localhost:5432/db"`
	GRPCPort     string `env:"GRPC_PORT" env-default:":50051"`
	ExchangeURL  string `env:"EXCHANGE_URL" env-default:"https://api.grinex.com/v1/rates"`
	LogLevel     string `env:"LOG_LEVEL" env-default:"dev"`
	OtelEndpoint string `env:"OTEL_EXPORTER_OTLP_ENDPOINT" env-default:"http://localhost:4317"`
}

func MustLoadConfig() *Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Println("failed to read .env file")
	}

	dbURL := flag.String("db-url", cfg.DatabaseURL, "Database connection string")
	grpcPort := flag.String("grpc-port", cfg.GRPCPort, "gRPC server port")
	flag.Parse()

	if isFlagPassed("db-url") {
		cfg.DatabaseURL = *dbURL
	}
	if isFlagPassed("grpc-port") {
		cfg.GRPCPort = *grpcPort
	}

	return &cfg
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
