package entity

import (
	"errors"

	"github.com/monkeydioude/drannoc/internal/misc"
)

type Prices map[string]float64

func NewPrices() Prices {
	return make(Prices)
}

func (p Prices) Set(currency string, amount float64) error {
	if !misc.AvailableCurrency(currency) {
		return errors.New("currency does not exist in config")
	}
	p[currency] = amount
	return nil
}

func (p Prices) IsEmpty() bool {
	return len(p) == 0
}
