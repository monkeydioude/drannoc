package service

import (
	"time"

	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/misc"
	repo "github.com/monkeydioude/drannoc/internal/repository"
)

// CreateAuthTokenNow is the auth token creation routine
func CreateAuthTokenNow(
	tokenRepo *repo.AuthToken,
) (*entity.AuthToken, error) {
	start := time.Now()
	duration := misc.TokenDuration
	token := entity.GenerateAuthToken(start, duration)
	_, err := tokenRepo.Store(token)
	return token, err
}
