package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/WhilsoM/test-go-senior/services/rate-service/internal/dto"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type RateClient interface {
	FetchOrderBook(ctx context.Context) ([]float64, []float64, error)
	Close() error
}

type rateClient struct {
	client *resty.Client
	url    string
	log    *zap.Logger
}

func NewRateClient(url string, log *zap.Logger) RateClient {
	return &rateClient{
		client: resty.New(),
		url:    url,
		log:    log,
	}
}

func (c *rateClient) Close() error {
	return c.client.Close()
}

// fetch data and convert strings to float64
func (c *rateClient) FetchOrderBook(ctx context.Context) ([]float64, []float64, error) {
	var result dto.OrderBookResponse
	c.log.Info("FetchOrderBook make a request to", zap.String("url", c.url))

	resp, err := c.client.R().
		SetContext(ctx).
		SetResult(&result).
		Get(c.url)
	if err != nil {
		c.log.Error("Error with a request", zap.Error(err))
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		c.log.Error("Api returned error", zap.Error(err))
		return nil, nil, fmt.Errorf("api returned error: %s", resp.Status())
	}

	asks := make([]float64, len(result.Asks))
	for i, item := range result.Asks {
		asks[i], _ = strconv.ParseFloat(item.Price, 64)
	}

	bids := make([]float64, len(result.Bids))
	for i, item := range result.Bids {
		bids[i], _ = strconv.ParseFloat(item.Price, 64)
	}

	c.log.Info("FetchOrderBook is done")

	return asks, bids, nil
}
