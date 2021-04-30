package entity

type Coin struct {
	ID     string             `json:"id"`
	Symbol string             `json:"symbol"`
	Prices map[string]float32 `json:"prices"`
}
