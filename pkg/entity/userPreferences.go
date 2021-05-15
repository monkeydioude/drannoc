package entity

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type UserPreferences struct {
	ID    string   `json:"id,omitempty"`
	Coins []string `json:"coins"`
}

func (p *UserPreferences) GetID() string {
	return p.ID
}

func (p *UserPreferences) SetID(id string) {
	p.ID = id
}

// String implements the Stringer interface
func (p *UserPreferences) String() string {
	res, err := json.Marshal(p)

	if err != nil {
		log.Error(err)
		return ""
	}
	return string(res)
}
