package blockchain

import "errors"

type Currency string

const (
	USD = "USD"
	AUD = "AUD"
	BRL = "BRL"
)

func (c Currency) IsValid() error {
	switch c {
	case USD, AUD, BRL:
		return nil
	}
	return errors.New("Unsupported currency")
}
