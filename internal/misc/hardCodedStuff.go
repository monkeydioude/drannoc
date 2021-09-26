package misc

import "time"

// this file should be gradually removed

const TokenDuration = 10 * 24 * time.Hour

var AvailableCurrencies = []string{"USD", "EUR"}

func AvailableCurrency(currency string) bool {
	for _, curr := range AvailableCurrencies {
		if curr == currency {
			return true
		}
	}
	return false
}
