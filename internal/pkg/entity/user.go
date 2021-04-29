package entity

import (
	"encoding/json"

	"github.com/monkeydioude/drannoc/internal/pkg/encrypt"
	log "github.com/sirupsen/logrus"
)

// User represents the data of a user
type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Created  int64  `json:"created"`
	Password string `json:"password"`
	// AuthID string `json:"authID"`
	Email string `json:"email"`
}

// GetID the bolt.Entity interface
func (u *User) GetID() string {
	return u.ID
}

// String implements the Stringer interface
func (u *User) String() string {
	res, err := json.Marshal(u)

	if err != nil {
		log.Error(err)
		return ""
	}
	return string(res)
}

// SetID implements the Stringer interface
func (u *User) SetID(id string) {
	u.ID = id
}

// SetID implements the Stringer interface
func (u *User) PasswordEncrypt() {
	u.Password = encrypt.MD5FromString(u.Password)
}

// NewUser creates a pointer to a new User instance
func NewUser(login, password string) (*User, error) {
	h := encrypt.MD5FromString(password)
	// if err != nil {
	// 	return nil, err
	// }
	return &User{
		Login:    login,
		Password: h,
	}, nil
}
