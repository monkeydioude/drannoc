package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/body"
	"github.com/monkeydioude/drannoc/internal/service"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// Authenticate handles any kind of authentification (user, bot etc...)
func Authenticate(c *gin.Context) {
	loginData, err := body.NewLoginData(c.Request.Body)
	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	if !loginData.IsValid() {
		res.Write(c, res.BadRequest("Could not retrieve login data"))
		return
	}

	authToken, err := service.Authenticate(loginData)

	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	res.Ok(c, gin.H{
		"token": authToken.GetToken(),
		"data":  authToken,
	})
}
