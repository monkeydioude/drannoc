package entity

import "math"

type Coin struct {
	ID     string             `json:"id"`
	Symbol string             `json:"symbol"`
	Prices map[string]float64 `json:"prices"`
}

func (c Coin) RoundPrices(limit int) {
	for key, p := range c.Prices {
		n := math.Pow10(limit)
		c.Prices[key] = math.Round(p*n) / n
	}
}
