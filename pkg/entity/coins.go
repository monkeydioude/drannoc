package entity

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Stack struct {
	ID         string          `json:"_id,omitempty"`
	Coins      map[string]Coin `json:"coins"`
	Created_at int64           `json:"created_at"`
}

func (c Stack) GetID() string {
	return c.ID
}

func (c Stack) SetID(_ string) {
}

func (c Stack) String() string {
	res, err := json.Marshal(c)

	if err != nil {
		log.Error(err)
		return ""
	}
	return string(res)
}

func (c Stack) RoundCoinsPrices(limit int) {
	for _, c := range c.Coins {
		c.RoundPrices(limit)
	}
}
