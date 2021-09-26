package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/misc"
	repo "github.com/monkeydioude/drannoc/internal/repository"
	"github.com/monkeydioude/drannoc/pkg/db"
)

// CreateAuthTokenNow is the auth token creation routine
func CreateAuthTokenNow(
	consumer string,
	tokenLives int,
) *entity.AuthToken {
	start := time.Now()
	duration := misc.TokenDuration
	token := entity.GenerateAuthToken(start, duration, consumer, tokenLives)
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
	tokenLives int,
) bool {
	if !token.ShouldRemakeNow() {
		return false
	}

	t := CreateAuthTokenNow(token.Consumer, tokenLives)
	*token = *t

	return true
}

type Identifiers struct {
	AuthToken string
	Consumer  string
}

func FindIdentifiers(c *gin.Context) (Identifiers, error) {
	authTokenLabel := c.GetString("AuthTokenLabel")
	consumerLabel := c.GetString("ConsumerLabel")
	authToken, _ := c.Cookie(authTokenLabel)
	consumer, _ := c.Cookie(consumerLabel)

	if authToken != "" && consumer != "" {
		return Identifiers{
			AuthToken: authToken,
			Consumer:  consumer,
		}, nil
	}

	authToken = c.GetHeader(authTokenLabel)
	consumer = c.GetHeader(consumerLabel)

	if authToken != "" && consumer != "" {
		return Identifiers{
			AuthToken: authToken,
			Consumer:  consumer,
		}, nil
	}

	return Identifiers{"", ""},
		errors.New("looked very hard but could not find identifiers")
}
