package service

import (
	"context"
	"fmt"

	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/client"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/dto"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

type RateService interface {
	GetRates(ctx context.Context, topN, avgN, avgM int) (dto.RateResult, error)
}

type rateService struct {
	client client.RateClient
	repo   repository.RateRepository
	log    *zap.Logger
}

func NewRateService(client client.RateClient, repo repository.RateRepository, log *zap.Logger) RateService {
	return &rateService{
		client: client,
		repo:   repo,
		log:    log,
	}
}

func (s *rateService) GetRates(ctx context.Context, topN, avgN, avgM int) (dto.RateResult, error) {
	tracer := otel.Tracer("rate-service")
	ctx, span := tracer.Start(ctx, "Service.GetRates")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()

	s.log.Info("GetRates started")

	asks, bids, err := s.client.FetchOrderBook(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "failed to fetch order book")
		span.RecordError(err)
		s.log.Error("GetRates failed to fetch order book", zap.String("trace_id", traceID), zap.Error(err))
		return dto.RateResult{}, fmt.Errorf("fetch error: %w", err)
	}

	res := dto.RateResult{
		AskTopN: calculateTopN(asks, topN),
		BidTopN: calculateTopN(bids, topN),
		AskAvg:  calculateAvg(asks, avgN, avgM),
		BidAvg:  calculateAvg(bids, avgN, avgM),
	}

	if err := s.repo.SaveRate(ctx, res.AskTopN, res.BidTopN); err != nil {
		span.SetStatus(codes.Error, "failed to save db")
		span.RecordError(err)
		s.log.Error("GetRates db save error", zap.Error(err))
	}

	return res, nil
}

// get value at specific index
func calculateTopN(data []float64, n int) float64 {
	if n < 0 || n >= len(data) {
		return 0
	}
	return data[n]
}

// calculate average between n and m
func calculateAvg(data []float64, n, m int) float64 {
	if n < 0 || m < 0 || n > m || m >= len(data) {
		return 0
	}

	var sum float64
	count := 0
	for i := n; i <= m; i++ {
		sum += data[i]
		count++
	}

	if count == 0 {
		return 0
	}
	return sum / float64(count)
}
