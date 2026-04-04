package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	ratesv1 "github.com/WhilsoM/test-go-senior/gen/go/rates/v1"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/client"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/config"
	grpcserver "github.com/WhilsoM/test-go-senior/services/rate-service/internal/grpc"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/repository"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/service"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/tracer"

	"github.com/WhilsoM/test-go-senior/core/logger"
	"github.com/WhilsoM/test-go-senior/core/storage"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Init config and logger
	cfg := config.MustLoadConfig()
	log := logger.NewLogger(cfg.LogLevel)
	defer log.Sync()

	ctx := context.Background()

	// Init tracer
	tp, err := tracer.InitTracer(ctx, "rate-service", cfg.OtelEndpoint)
	if err != nil {
		log.Fatal("failed to init tracer", zap.Error(err))
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Error("failed to shutdown tracer", zap.Error(err))
		}
	}()

	// Init db
	pool := storage.MustLoadDatabase(cfg.DatabaseURL, "./migrations")
	defer pool.Close()

	// Init layers
	repo := repository.NewRateRepository(pool, log)
	rateClient := client.NewRateClient(cfg.ExchangeURL, log)
	defer rateClient.Close()

	ratesService := service.NewRateService(rateClient, repo, log)

	ratesHandler := grpcserver.NewServer(ratesService, log)

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	ratesv1.RegisterRatesServiceServer(grpcServer, ratesHandler)

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}

	go func() {
		log.Info("gRPC server started", zap.String("port", cfg.GRPCPort))
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to serve", zap.Error(err))
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Info("shutting down gracefully...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	log.Info("server stopped")
}
