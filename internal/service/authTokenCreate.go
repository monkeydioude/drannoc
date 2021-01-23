package service

import (
	"time"

	"github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/entity"
)

// AuthTokenCreate is the auth token creation routine
func AuthTokenCreate(password string, start time.Time, duration time.Duration) (*entity.AuthToken, error) {
	token := entity.NewAuthToken(password, start, duration)
	_, err := bucket.AuthToken(nil).Store(token)
	return token, err
}
