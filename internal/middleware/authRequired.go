package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/config"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	"github.com/monkeydioude/drannoc/internal/routine"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// AuthRequired is a middleware for checking
// token authentifier presence and availability
func AuthRequired(c *gin.Context) {
	// fetching token from Header
	tokenID := c.GetHeader(config.AuthTokenLabel)
	if tokenID == "" {
		res.Write(c, res.Redirect(config.UserLoginRoute))
		return
	}

	tokenRepo := repository.NewAuthToken()
	token := &entity.AuthToken{
		Token: tokenID,
	}

	// loading token from DB
	_, err := tokenRepo.Load(token)
	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	// check token validity
	if token == nil || token.GetToken() != tokenID || !token.IsValidNow() {
		c.SetCookie(config.AuthTokenLabel, "", -1, "/", "", false, false)
		res.Write(c, res.Redirect(config.UserLoginRoute))
		return
	}

	// trying to regenerate the token
	didItRegenerate, err := routine.TryRegenerateToken(
		repository.NewAuthToken(),
		token,
	)

	if err != nil {
		res.Write(c, res.ServiceUnavailable("auth-token refresh issue", err.Error()))
		return
	}

	// if token was regenerated, then refresh auth headers as well
	if didItRegenerate {
		// resetting cookie
		// @todo refactorize-token-consumer-setcookie
		maxAge := int(token.Expires - time.Now().Unix())
		c.SetCookie(config.AuthTokenLabel, token.GetToken(), maxAge, "/", "", false, false)
		c.SetCookie(config.ConsumerLabel, token.Consumer, maxAge, "/", "", false, false)
		return
	}

	// not regenerated, token tick once, applying tick altearations
	// (removing a life of the token, for example)
	token.Tick()

	// then storing it
	err = tokenRepo.Save(token)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("auth-token refresh issue", err.Error()))
		return
	}
}
