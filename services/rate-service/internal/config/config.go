package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `env:"ENV" env-default:"dev"`
	GRPC        string `env:"GRPC" env-default:":50051"`
	DatabaseURL string `env:"DATABASE_URL" env-required:"true"`
	ExchangeURL string `env:"EXCHANGE_URL" env-required:"true"`
}

func MustLoadConfig() *Config {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return &cfg
}
