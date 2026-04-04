package grpc

import (
	"context"

	ratesv1 "github.com/WhilsoM/test-go-senior/gen/go/rates/v1"
	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/service"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RateServer struct {
	ratesv1.UnimplementedRatesServiceServer
	svc service.RateService
	log *zap.Logger
}

func NewServer(svc service.RateService, log *zap.Logger) *RateServer {
	return &RateServer{svc: svc, log: log}
}

func (s *RateServer) GetRates(ctx context.Context, req *ratesv1.GetRatesRequest) (*ratesv1.GetRatesResponse, error) {
	s.log.Info("GetRates started")

	res, err := s.svc.GetRates(ctx, int(req.TopN), int(req.AvgN), int(req.AvgM))
	if err != nil {
		s.log.Error("GetRates have an error", zap.Error(err))
		return nil, err
	}

	s.log.Info("GetRates is done")

	return &ratesv1.GetRatesResponse{
		AskTopN:   res.AskTopN,
		BidTopN:   res.BidTopN,
		AskAvgNm:  res.AskAvg,
		BidAvgNm:  res.BidAvg,
		CreatedAt: timestamppb.Now(),
	}, nil
}

func (s *RateServer) HealthCheck(ctx context.Context, req *ratesv1.HealthCheckRequest) (*ratesv1.HealthCheckResponse, error) {
	s.log.Info("HealthCheck")
	return &ratesv1.HealthCheckResponse{Status: "OK"}, nil
}
