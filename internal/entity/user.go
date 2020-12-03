package entity

import (
	"encoding/json"
	"fmt"

	"github.com/monkeydioude/drannoc/internal/bucket"
	log "github.com/sirupsen/logrus"
)

// User represents the data of a user
type User struct {
	id     string
	AuthID string `json:"authID"`
	Email  string `json:"email"`
}

// GetKey the bolt.Entity interface
func (u *User) GetKey() string {
	return u.id
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

// LoadUser try to load an existing from DB
func LoadUser(id string) (*User, error) {
	userBucket := bucket.User(nil)

	u, err := userBucket.Get(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("Could not find user %s", id)
	}

	user := &User{
		id: id,
	}

	return user, json.Unmarshal(u, user)
}
