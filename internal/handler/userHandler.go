package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/body"
	"github.com/monkeydioude/drannoc/internal/service"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

const (
	loginKey    string = "l"
	passwordKey string = "p"
)

// UserCreate handles user creation form
// POST /user
func UserCreate(c *gin.Context) {
	loginData, err := body.NewLoginData(c.Request.Body)
	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	token, err := service.UserCreate(loginData)

	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	if token == nil {
		res.ServiceUnavailable("Could not create user", "Could not create user")
		return
	}

	res.Ok(c, gin.H{
		"token": token.GetToken(),
		"data":  token,
	})
}

// UserLogin handles user login form
// POST /user/login
func UserLogin(c *gin.Context) {
	// loginData, err := body.NewLoginData(c.Request.Body)
	// if err != nil {
	// 	res.Write(c, res.BadRequest(err.Error()))
	// 	return
	// }
	// userBucket := bucket.User(nil)

	// u, err := userBucket.Get(loginData.Login)
	// if err != nil {
	// 	res.Write(c, res.ServiceUnavailable("Could not retrieve user", err.Error()))
	// 	return
	// }

	// if u == nil {
	// 	res.Write(c, res.BadRequest(fmt.Sprintf("Username `%s` does not exist", loginData.Login)))
	// 	return
	// }

	// authToken, err := service.Authenticate(loginData)

	// if err != nil {
	// 	res.Write(c, res.ServiceUnavailable("Could not authenticate", err.Error()))
	// 	return
	// }

	// res.Ok(c, gin.H{
	// 	"token": authToken.GetToken(),
	// 	"data":  authToken,
	// })
}
