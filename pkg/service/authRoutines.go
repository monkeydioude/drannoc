package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/config"
	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/misc"
	repo "github.com/monkeydioude/drannoc/pkg/repository"
)

// CreateAuthTokenNow is the auth token creation routine
func CreateAuthTokenNow(
	consumer string,
) *entity.AuthToken {
	start := time.Now()
	duration := misc.TokenDuration
	token := entity.GenerateAuthToken(start, duration, consumer)
	return token
}

// RevokeAuthToken is called when want to render a token unusable.
// e.g.: login out
func RevokeAuthToken(tokenRepo *repo.AuthToken, token string) error {
	expires := time.Now().Unix()

	entity := &entity.AuthToken{
		Token: token,
	}
	_, err := tokenRepo.FindFirst(entity, db.Filter{"token": token}, nil)
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
	token *entity.AuthToken,
) bool {
	if !token.ShouldRemakeNow() {
		return false
	}

	t := CreateAuthTokenNow(token.Consumer)
	*token = *t

	return true
}

type Identifiers struct {
	AuthToken string
	Consumer  string
}

func FindIdentifiers(c *gin.Context) (Identifiers, error) {
	authToken, _ := c.Cookie(config.AuthTokenLabel)
	consumer, _ := c.Cookie(config.ConsumerLabel)

	if authToken != "" && consumer != "" {
		return Identifiers{
			AuthToken: authToken,
			Consumer:  consumer,
		}, nil
	}

	authToken = c.GetHeader(config.AuthTokenLabel)
	consumer = c.GetHeader(config.ConsumerLabel)

	if authToken != "" && consumer != "" {
		return Identifiers{
			AuthToken: authToken,
			Consumer:  consumer,
		}, nil
	}

	return Identifiers{"", ""},
		errors.New("looked very hard but could not find identifiers")
}
