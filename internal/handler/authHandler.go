package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/config"
	"github.com/monkeydioude/drannoc/internal/repository"
	"github.com/monkeydioude/drannoc/internal/routine"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// UserLogout handles user loging out
// Delete /auth
func AuthDelete(c *gin.Context) {
	token := c.GetHeader(config.AuthTokenLabel)

	// revoking token if exist in header
	if token != "" {
		routine.RevokeAuthToken(repository.NewAuthToken(), token)
	}

	// unsetting cookies, since user asked for logout
	c.SetCookie(config.AuthTokenLabel, "", -1, "/", "", false, false)
	c.SetCookie(config.ConsumerLabel, "", -1, "/", "", false, false)
	// OKKKK
	res.Ok(c, gin.H{})
}
