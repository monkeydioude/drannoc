package entity

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type CoinInfo struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	Created_at int64  `json:"created_at"`
}

// GetID the bolt.Entity interface
func (c *CoinInfo) GetID() string {
	return c.ID
}

// String implements the Stringer interface
func (c *CoinInfo) String() string {
	res, err := json.Marshal(c)

	if err != nil {
		log.Error(err)
		return ""
	}
	return string(res)
}

func (c *CoinInfo) SetID(id string) {
	c.ID = id
}
