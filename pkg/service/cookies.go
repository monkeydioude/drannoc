package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/config"
	"github.com/monkeydioude/drannoc/pkg/entity"
)

func SetCookies(token *entity.AuthToken, user string, c *gin.Context) {
	maxAge := int(token.Expires - time.Now().Unix())
	// setting up the AuthToken Cookie
	c.SetCookie(config.AuthTokenLabel, token.GetToken(), maxAge, "/", "", false, false)
	c.SetCookie(config.ConsumerLabel, user, maxAge, "/", "", false, false)
}

func UnsetCookies(c *gin.Context) {
	c.SetCookie(config.AuthTokenLabel, "", -1, "/", "", false, false)
	c.SetCookie(config.ConsumerLabel, "", -1, "/", "", false, false)
}
