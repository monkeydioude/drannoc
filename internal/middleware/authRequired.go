package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// AuthRequired is a middleware for checking
// token authentifier presence and availability
func AuthRequired(c *gin.Context) {
	tokenID := c.GetHeader("auth-token")

	if tokenID == "" {
		res.Redirect(c, "/login")
		return
	}
	token, err := entity.LoadAuthToken(tokenID)
	fmt.Println(token.GetToken())

	if err != nil {
		res.BadRequest(c, err.Error())
		return
	}

	if token == nil || token.GetToken() != tokenID || !token.IsValidNow() {
		res.Redirect(c, "/login")
		return
	}
}
