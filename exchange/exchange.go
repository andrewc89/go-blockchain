package exchange

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
	defer resp.Body.Close()
	tickers := make(map[string]CurrencyTicker)
	jsonErr := json.NewDecoder(resp.Body).Decode(&tickers)
	if jsonErr != nil {
		return nil, jsonErr
	}
	ct := tickers[currency]
	return &ct, nil
}
