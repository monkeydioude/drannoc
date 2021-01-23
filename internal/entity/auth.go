package entity

import (
	"time"

	"github.com/monkeydioude/drannoc/internal/encrypt"
)

// Auth matches a `auth`'s bucket entity
type Auth struct {
	ID       string `json:"id"`
	UserID   string `json:"userID"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Created  int    `json:"created"`
}

// GetID implements the entity.Entity interface
func (a *Auth) GetID() string {
	return a.ID
}

// String implements the Stringer interface
func (a *Auth) String() string {
	return a.Password
}

// GetPassword returns en
func (a *Auth) GetPassword() string {
	return string(a.Password)
}

// SetID implements the Stringer interface
func (a *Auth) SetID(id string) {
	a.ID = id
}

// NewAuth generates a new Auth *struct
func NewAuth(userID, password string) *Auth {
	return &Auth{
		UserID:   userID,
		Password: encrypt.Password(userID, password),
		Created:  time.Now().Second(),
	}
}
