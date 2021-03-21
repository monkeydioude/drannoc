package routine

import (
	"time"

	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/misc"
	repo "github.com/monkeydioude/drannoc/internal/repository"
)

// CreateAuthTokenNow is the auth token creation routine
func CreateAuthTokenNow(
	tokenRepo *repo.AuthToken,
	consumerID string,
) (*entity.AuthToken, error) {
	start := time.Now()
	duration := misc.TokenDuration
	token := entity.GenerateAuthToken(start, duration, consumerID)
	_, err := tokenRepo.Store(token)
	return token, err
}

// RevokeAuthToken is called when want to render a token unusable.
// e.g.: login out
func RevokeAuthToken(tokenRepo *repo.AuthToken, token string) error {
	expires := time.Now().Unix()

	entity := &entity.AuthToken{
		Token: token,
	}
	_, err := tokenRepo.FindFirst(entity, repo.Filter{"token": token})
	if err != nil {
		return err
	}

	entity.Expires = expires

	_, err = tokenRepo.Store(entity)
	if err != nil {
		return err
	}

	return nil
}
