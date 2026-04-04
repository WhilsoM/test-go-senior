package dto

type OrderItem struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
}

type OrderBookResponse struct {
	Asks []OrderItem `json:"asks"`
	Bids []OrderItem `json:"bids"`
}

type RateResult struct {
	AskTopN float64
	BidTopN float64
	AskAvg  float64
	BidAvg  float64
}
