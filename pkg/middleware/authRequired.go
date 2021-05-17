package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/config"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	routine "github.com/monkeydioude/drannoc/pkg/service"
)

// AuthRequired is a middleware for checking
// token authentifier presence and availability
func AuthRequired(c *gin.Context) {
	// fetching token from Header
	ids, err := routine.FindIdentifiers(c)

	if err != nil {
		c.Abort()
		res.Write(c, res.Unauthorized(config.UserLoginRoute))
		return
	}

	tokenRepo := repository.NewAuthToken()
	token := &entity.AuthToken{
		Token: ids.AuthToken,
	}

	// loading token from DB
	_, err = tokenRepo.Load(token)
	if err != nil {
		c.Abort()
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	// check token validity
	if token == nil ||
		token.GetToken() != ids.AuthToken ||
		!token.IsValidNow() ||
		token.Consumer != ids.Consumer {
		routine.UnsetCookies(c)
		res.Write(c, res.Unauthorized(config.UserLoginRoute))
		c.Abort()
		return
	}

	// trying to regenerate the token
	didItRegenerate := routine.TryRegenerateToken(token)

	// token tick once, applying tick alterations
	// (removing a life of the token, for example)
	token.Tick()

	// if token was regenerated, then refresh auth headers as well
	if didItRegenerate {
		// resetting cookie
		// @todo refactorize-token-consumer-setcookie
		routine.SetCookies(token, ids.Consumer, c)
		// then saving it
		_, err = tokenRepo.Store(token)
	} else {
		// then storing it
		err = tokenRepo.Save(token)
	}

	if err != nil {
		res.Write(c, res.ServiceUnavailable("auth-token refresh issue", err.Error()))
		c.Abort()
		return
	}

	// setting identifiers in context for further use
	c.Set("auth-token", ids.AuthToken)
	c.Set("consumer", ids.Consumer)
}
