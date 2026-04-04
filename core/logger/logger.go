package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init logger dev and prod modes
func NewLogger(env string) *zap.Logger {
	var config zap.Config

	if env == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.OutputPaths = []string{"stdout"}

	logger, err := config.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	return logger
}
