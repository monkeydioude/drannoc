package entity

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/encrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// AuthToken represents a bearer token entity. One should pass it
// along any request on private endpoints
type AuthToken struct {
	Created  int64 `json:"created"`
	Expires  int64 `json:"expires"`
	LastUsed int64 `json:"lastUsed"`
	token    string
}

// GetKey the bolt.Entity interface
func (a *AuthToken) GetKey() string {
	return a.token
}

// GetToken just gets the token aka the key
func (a *AuthToken) GetToken() string {
	return a.GetKey()
}

// String implements the Stringer interface
func (a *AuthToken) String() string {
	data, err := json.Marshal(a)

	if err != nil {
		return ""
	}
	return string(data)
}

// IsValid verifies a token is still valid
func (a *AuthToken) IsValid(date time.Time) bool {
	return date.Before(time.Unix(a.Created, 0))
}

// IsValidNow is the same as IsValid, but now
func (a *AuthToken) IsValidNow() bool {
	return a.IsValid(time.Now())
}

// LoadAuthToken retrieve
func LoadAuthToken(token string) (*AuthToken, error) {
	bucket := bucket.AuthToken(nil)

	data, err := bucket.Get(token)
	if err != nil {
		return nil, err
	}
	authToken := &AuthToken{
		token: token,
	}

	return authToken, json.Unmarshal(data, authToken)
}

// NewAuthToken generates a new Authentification token along
// with its time related data
func NewAuthToken(passwd string, duration time.Duration, date time.Time) *AuthToken {
	const len int = 12
	var buffer bytes.Buffer
	created := date.Unix()
	lastUsed := created
	expires := date.Add(duration).Unix()

	for i := 0; i < len; i++ {
		buffer.WriteByte(byte(rand.Intn(50)))
	}

	buffer.WriteString(passwd)
	return &AuthToken{
		Created:  created,
		Expires:  expires,
		LastUsed: lastUsed,
		token:    encrypt.MD5FromBytes(buffer.Bytes()),
	}
}
