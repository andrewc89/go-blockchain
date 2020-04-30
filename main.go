package main

import (
	"fmt"

	"github.com/gwely/blockchain/exchange"
)

func main() {
	exchange := exchange.NewBlockchainExchange()
	ticker, err := exchange.Ticker("USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(ticker)
}
