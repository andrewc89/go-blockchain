package main

import (
	"fmt"

	"github.com/gwely/go-blockchain/blockchain"
)

func main() {
	exchange := blockchain.NewBlockchainExchange()
	ticker, err := exchange.Ticker("USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(ticker)
}
