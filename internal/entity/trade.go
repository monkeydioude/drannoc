package entity

import (
	"encoding/json"
	"time"

	"github.com/monkeydioude/drannoc/pkg/slice"
	log "github.com/sirupsen/logrus"
)

var directions = slice.String{"BUY", "SELL"}

type Trade struct {
	ID          string  `json:"id"`
	Parent_id   string  `json:"parent_id"`
	User_id     string  `json:"user_id,omitempty"`
	Coin_id     string  `json:"coin_id"`
	Created_at  int64   `json:"created_at"`
	Traded_at   int64   `json:"traded_at"`
	Modified_at int64   `json:"modified_at"`
	Prices      Prices  `json:"prices"`
	Direction   string  `json:"direction"`
	Amount      float64 `json:"amount"`
}

// GetID implements Entity interface
func (t *Trade) GetID() string {
	return t.ID
}

// SetID implements Entity interface
func (t *Trade) SetID(id string) {
	t.ID = id
}

// String implements Entity interface
func (t *Trade) String() string {
	res, err := json.Marshal(t)

	if err != nil {
		log.Error(err)
		return ""
	}
	return string(res)
}

func (t *Trade) IsValidDirection() bool {
	for _, d := range directions {
		if d == t.Direction {
			return true
		}
	}
	return false
}

func (t *Trade) IsStorable() bool {
	return t.User_id != "" &&
		t.Created_at != 0 &&
		t.Modified_at != 0 &&
		t.Traded_at != 0 &&
		!t.Prices.IsEmpty() &&
		t.IsValidDirection() &&
		directions.Contains(t.Direction)
}

func (t *Trade) UpdateWith(mergeOver *Trade) {
	if mergeOver.Coin_id != "" {
		t.Coin_id = mergeOver.Coin_id
	}

	if len(mergeOver.Prices) > 0 {
		t.Prices = mergeOver.Prices
	}

	if mergeOver.Direction != "" {
		t.Direction = mergeOver.Direction
	}

	if mergeOver.Amount > 0 {
		t.Amount = mergeOver.Amount
	}

	if mergeOver.Traded_at > 0 {
		t.Traded_at = mergeOver.Traded_at
	}

	t.Modified_at = time.Now().UnixNano()
}

func NewTrade() *Trade {
	return &Trade{}
}
