package blockchain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlockchainExchange struct {
	apiCode *string
}

type CurrencyTicker struct {
	FifteenMinutes float32 `json:"15m"`
	Last           float32 `json:"last"`
	Buy            float32 `json:"buy"`
	Sell           float32 `json:"sell"`
	Symbol         string  `json:"symbol"`
}

type ApiError struct {
	Status string
	Url    string
}

func (err ApiError) Error() string {
	return fmt.Sprintf("Received %s from %s", err.Status, err.Url)
}

func NewBlockchainExchange() *BlockchainExchange {
	return &BlockchainExchange{}
}

func NewBlockchainExchangeWithAuth(apiCode *string) *BlockchainExchange {
	return &BlockchainExchange{
		apiCode: apiCode,
	}
}

const endpoint = "https://blockchain.info/ticker"

func (e *BlockchainExchange) Ticker(currency string) (*CurrencyTicker, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, ApiError{resp.Status, endpoint}
	}
	defer resp.Body.Close()
	tickers := make(map[string]CurrencyTicker)
	jsonErr := json.NewDecoder(resp.Body).Decode(&tickers)
	if jsonErr != nil {
		return nil, jsonErr
	}
	ct := tickers[currency]
	return &ct, nil
}
