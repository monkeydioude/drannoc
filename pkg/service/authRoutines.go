package service

import (
	"time"

	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/misc"
	repo "github.com/monkeydioude/drannoc/pkg/repository"
)

// CreateAuthTokenNow is the auth token creation routine
func CreateAuthTokenNow(
	tokenRepo *repo.AuthToken,
	consumer string,
) (*entity.AuthToken, error) {
	start := time.Now()
	duration := misc.TokenDuration
	token := entity.GenerateAuthToken(start, duration, consumer)
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
	_, err := tokenRepo.FindFirst(entity, db.Filter{"token": token})
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

// TryRegenerateToken try to generate a new token. First check
// if we should regenerate a new
func TryRegenerateToken(
	tokenRepo *repo.AuthToken,
	token *entity.AuthToken,
) (bool, error) {
	if !token.ShouldRemakeNow() {
		return false, nil
	}

	repo := repo.NewAuthToken()
	t, err := CreateAuthTokenNow(
		repo,
		token.Consumer,
	)
	*token = *t

	if err != nil {
		return false, err
	}

	return true, nil
}
