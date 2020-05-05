package blockchain

import (
	"encoding/json"
	"net/http"
)

type CurrencyTicker struct {
	FifteenMinutes float32 `json:"15m"`
	Last           float32 `json:"last"`
	Buy            float32 `json:"buy"`
	Sell           float32 `json:"sell"`
	Symbol         string  `json:"symbol"`
}

const tickerUrl = domain + "/ticker"

type tickerQueryParameters struct {
	apiCode *string
}

func (qps tickerQueryParameters) QueryParameters() QueryParameters {
	qpMap := make(map[string]string)
	if qps.apiCode != nil {
		qpMap["apiCode"] = *qps.apiCode
	}
	return QueryParameters(qpMap)
}

func (exchange *BlockchainExchange) Ticker(currency Currency) (*CurrencyTicker, error) {
	if currencyErr := currency.IsValid(); currencyErr != nil {
		return nil, currencyErr
	}

	qps := tickerQueryParameters{exchange.apiCode}
	endpoint := Endpoint{tickerUrl, qps.QueryParameters()}

	resp, httpErr := http.Get(endpoint.String())
	if httpErr != nil {
		return nil, httpErr
	}
	if resp.Status != statusOk {
		return nil, ApiError{resp.Status, endpoint.String()}
	}
	defer resp.Body.Close()

	tickers := make(map[string]CurrencyTicker)
	if jsonErr := json.NewDecoder(resp.Body).Decode(&tickers); jsonErr != nil {
		return nil, jsonErr
	}

	ct := tickers[string(currency)]
	return &ct, nil
}
