package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/config"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// AuthRequired is a middleware for checking
// token authentifier presence and availability
func AuthRequired(c *gin.Context) {
	tokenID := c.GetHeader(config.AuthTokenLabel)
	if tokenID == "" {
		res.Write(c, res.Redirect(config.UserLoginRoute))
		return
	}

	token := &entity.AuthToken{
		Token: tokenID,
	}

	_, err := repository.NewAuthToken().Load(token)

	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	if token == nil || token.GetToken() != tokenID || !token.IsValidNow() {
		c.SetCookie(config.AuthTokenLabel, "", 0, "/", "", false, false)
		res.Write(c, res.Redirect(config.UserLoginRoute))
		return
	}
}
