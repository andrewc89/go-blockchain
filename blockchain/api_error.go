package blockchain

import (
  "fmt"
)

type ApiError struct {
	Status string
	Url    string
}

func (err ApiError) Error() string {
	return fmt.Sprintf("Received %s from %s", err.Status, err.Url)
}
