package blockchain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const toBtcUrl = domain + "/tobtc"

type toBtcQueryParameters struct {
	apiCode  *string
	currency *Currency
	value    *int
}

func (qps toBtcQueryParameters) QueryParameters() QueryParameters {
	qpMap := make(map[string]string)
	valStr := fmt.Sprintf("%d", *qps.value)
	if qps.apiCode != nil {
		qpMap["apiCode"] = *qps.apiCode
	}
	qpMap["currency"] = string(*qps.currency)
	qpMap["value"] = valStr
	return QueryParameters(qpMap)
}

func (exchange *BlockchainExchange) ToBtc(currency Currency, value int) (float64, error) {
	if currencyErr := currency.IsValid(); currencyErr != nil {
		return 0, currencyErr
	}

	qps := toBtcQueryParameters{exchange.apiCode, &currency, &value}
	endpoint := Endpoint{toBtcUrl, qps.QueryParameters()}

	resp, httpErr := http.Get(endpoint.String())
	if httpErr != nil {
		return 0, httpErr
	}
	if resp.Status != statusOk {
		return 0, ApiError{resp.Status, endpoint.String()}
	}
	defer resp.Body.Close()

	btcValueStr, respErr := ioutil.ReadAll(resp.Body)
	if respErr != nil {
		return 0, respErr
	}

	btcValue, parseErr := strconv.ParseFloat(string(btcValueStr), 64)
	if parseErr != nil {
		return 0, parseErr
	}
	return btcValue, nil
}
