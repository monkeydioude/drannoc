package entity

import (
	"github.com/monkeydioude/drannoc/internal/encrypt"
)

// Auth matches a `auth`'s bucket entity
type Auth struct {
	login    string
	password string
}

// GetKey the bolt.Entity interface
func (a *Auth) GetKey() string {
	return a.login
}

// String implements the Stringer interface
func (a *Auth) String() string {
	return a.password
}

// GetPassword returns en
func (a *Auth) GetPassword() string {
	return string(a.password)
}

// NewAuth generates a new Auth *struct
func NewAuth(login, password string) *Auth {
	return &Auth{
		login:    login,
		password: encrypt.Password(login, password),
	}
}
