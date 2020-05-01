package blockchain

import (
	"encoding/json"
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

const path = "https://blockchain.info/ticker"

func (exchange *BlockchainExchange) Ticker(currency string) (*CurrencyTicker, error) {
	qps := make(map[string]string)
	if exchange.apiCode != nil {
		qps["apiCode"] = *exchange.apiCode
	}
	endpoint := Endpoint{path, QueryParameters(qps)}
	resp, err := http.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, ApiError{resp.Status, endpoint.String()}
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
