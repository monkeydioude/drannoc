package entity

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/monkeydioude/drannoc/internal/config"
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
	Duration int   `json:"duration"`
	// ID can change but Token can persist
	// through time. For example if
	// a token is renewed
	Token    string `json:"token"`
	ID       string `json:"id"`
	Consumer string `json:"consumer"`
	Life     int    `json:"life"`
}

// GetID the bolt.Entity interface
func (a *AuthToken) GetID() string {
	return a.ID
}

// SetID the bolt.Entity interface
func (a *AuthToken) SetID(ID string) {
	a.ID = ID
}

// GetToken just gets the token aka the key
func (a *AuthToken) GetToken() string {
	return a.Token
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

// Tick remove 1 life from the total count
func (a *AuthToken) Tick() int {
	a.Life -= 1
	return a.Life
}

func (a *AuthToken) HasLife() bool {
	return a.Life > 0
}

// ShouldRemake verifies if the token shouldnt be re-generated
// before expiration. Expires * tokenRemakeThreshold < remake < Expires
func (a *AuthToken) ShouldRemake(date time.Time) bool {
	if !a.IsValid(date) {
		return false
	}

	t := int64(float64(time.Unix(a.Expires, 0).Unix()) * tokenRemakeThreshold)

	if date.After(time.Unix(t, 0)) && a.HasLife() {
		return true
	}

	return !a.HasLife()
}

// ShouldRemakeNow is the same as ShouldRemake, but now
func (a *AuthToken) ShouldRemakeNow() bool {
	return a.ShouldRemake(time.Now())
}

// GenerateAuthToken generates a new Authentification token along
// with its time related data
func GenerateAuthToken(
	start time.Time,
	duration time.Duration,
	consumer string,
) *AuthToken {
	const len int = 12
	var buffer bytes.Buffer
	created := start.Unix()
	lastUsed := created
	expires := start.Add(duration).Unix()

	for i := 0; i < len; i++ {
		buffer.WriteByte(byte(rand.Intn(50)))
	}

	return &AuthToken{
		Created:  created,
		Expires:  expires,
		LastUsed: lastUsed,
		Duration: int(duration.Seconds()),
		Token:    encrypt.MD5(buffer.Bytes()),
		Consumer: consumer,
		Life:     config.TokenLivesMaxAmount,
	}
}
