package entity

import (
	"testing"
)

func TestICanRoundPricesUpTo2FloatingNumbers(t *testing.T) {
	coin := Coin{
		"dummy_coin",
		"DCON",
		map[string]float64{
			"C1": 1.82998989,
			"C2": 1.2,
			"C3": 1.33,
			"C4": 1.444,
		},
	}

	coin.RoundPrices(2)

	if coin.Prices["C1"] != 1.83 {
		t.Fatalf("C1 should be equal to 1.83")
	}
	if coin.Prices["C2"] != 1.2 {
		t.Fatalf("C2 should be equal to 1.2")
	}
	if coin.Prices["C3"] != 1.33 {
		t.Fatalf("C3 should be equal to 1.33")
	}
	if coin.Prices["C4"] != 1.44 {
		t.Fatalf("C4 should be equal to 1.44")
	}
}

func TestICanRoundPricesUpTo5FloatingNumbers(t *testing.T) {
	coin := Coin{
		"dummy_coin",
		"DCON",
		map[string]float64{
			"C1": 1.010205,
			"C2": 1.2,
			"C3": 1.33333,
			"C4": 1.444444,
		},
	}

	coin.RoundPrices(5)

	if coin.Prices["C1"] != 1.01021 {
		t.Fatalf("C1 should be equal to 1.01021")
	}
	if coin.Prices["C2"] != 1.2 {
		t.Fatalf("C2 should be equal to 1.2")
	}
	if coin.Prices["C3"] != 1.33333 {
		t.Fatalf("C3 should be equal to 1.33333")
	}
	if coin.Prices["C4"] != 1.44444 {
		t.Fatalf("C4 should be equal to 1.44444")
	}
}
