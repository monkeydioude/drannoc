package entity

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/encrypt"
)

const tokenRemakeThreshold = 0.8

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

// GetID the bolt.Entity interface
func (a *AuthToken) GetID() string {
	return a.token
}

// SetID the bolt.Entity interface
func (a *AuthToken) SetID(token string) {
	a.token = token
}

// GetToken just gets the token aka the key
func (a *AuthToken) GetToken() string {
	return a.GetID()
}

// String implements the Stringer interface
func (a *AuthToken) String() string {
	data, err := json.Marshal(a)

	if err != nil {
		return ""
	}
	return string(data)
}

// Update proceeds to token updates. Should be call before storing it
func (a *AuthToken) Update() {
	a.LastUsed = time.Now().Unix()
}

// IsValid verifies a token is still valid
func (a *AuthToken) IsValid(date time.Time) bool {
	return date.Before(time.Unix(a.Expires, 0))
}

// IsValidNow is the same as IsValid, but now
func (a *AuthToken) IsValidNow() bool {
	return a.IsValid(time.Now())
}

// ShouldRemake verifies if the token shouldnt be re-generated
// before expiration. Expires * tokenRemakeThreshold < remake < Expires
func (a *AuthToken) ShouldRemake(date time.Time) bool {
	if !a.IsValid(date) {
		return false
	}

	t := int64(float64(time.Unix(a.Expires, 0).Unix()) * tokenRemakeThreshold)

	return date.After(time.Unix(t, 0))
}

// ShouldRemakeNow is the same as ShouldRemake, but now
func (a *AuthToken) ShouldRemakeNow() bool {
	return a.ShouldRemake(time.Now())
}

// LoadAuthToken retrieve
func LoadAuthToken(tokenID string) (*AuthToken, error) {
	bucket := bucket.AuthToken(nil)

	data, err := bucket.Get(tokenID)
	if err != nil {
		return nil, err
	}

	authToken := &AuthToken{
		token: tokenID,
	}
	return authToken, json.Unmarshal(data, authToken)
}

// NewAuthToken generates a new Authentification token along
// with its time related data
func NewAuthToken(passwd string, start time.Time, duration time.Duration) *AuthToken {
	const len int = 12
	var buffer bytes.Buffer
	created := start.Unix()
	lastUsed := created
	expires := start.Add(duration).Unix()

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
