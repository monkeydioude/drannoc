package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/entity"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

const (
	loginKey    string = "l"
	passwordKey string = "p"
)

// CreateUser handles user creation form (PUT method)
func CreateUser(c *gin.Context) {
	login := c.PostForm(loginKey)
	password := c.PostForm(passwordKey)
	authBucket := bucket.Auth(nil)
	authTokenBucket := bucket.AuthToken(nil)

	if u, _ := authBucket.Get(login); u != nil {
		res.BadRequest(c, fmt.Sprintf("Username `%s` already exists\n", login))
		return
	}

	auth := entity.NewAuth(login, password)
	_, err := authBucket.Store(auth)
	if err != nil {
		res.BadRequest(c, err.Error())
		return
	}

	token := entity.NewAuthToken(auth.GetPassword(), tokenDuration, time.Now())
	_, err = authTokenBucket.Store(token)

	if err != nil {
		res.ServiceUnavailable(c, fmt.Sprintf("Could not store authToken for user `%s`\n", login))
		return
	}

	res.Ok(c, gin.H{
		"token": token.GetToken(),
		"data":  token,
	})
}

// UserLogin handles user login form (POST method)
func UserLogin(c *gin.Context) {
	userBucket := bucket.User(nil)
	login := c.PostForm(loginKey)

	u, err := userBucket.Get(login)
	if err != nil {
		res.BadRequest(c, err.Error())
		return
	}

	if u == nil {
		res.BadRequest(c, fmt.Sprintf("Username `%s` does not exist", login))
		return
	}
	Authenticate(c)
}
