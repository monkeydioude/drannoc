package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/entity"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// Authenticate handles any kind of authentification (user, bot etc...)
func Authenticate(c *gin.Context) {
	login := c.PostForm(loginKey)
	password := c.PostForm(passwordKey)
	authBucket := bucket.Auth(nil)
	authTokenBucket := bucket.AuthToken(nil)

	u, err := authBucket.Get(login)
	if err != nil {
		res.BadRequest(c, err.Error())
		return
	}

	if u == nil {
		res.BadRequest(c, fmt.Sprintf("User `%s` does not exist\n", login))
		return
	}

	encPasswd := entity.NewAuth(login, password).GetPassword()
	if string(u) != encPasswd {
		res.BadRequest(c, fmt.Sprintf("Wrong password for `%s`\n", login))
		return
	}

	token := entity.NewAuthToken(encPasswd, time.Now(), tokenDuration)
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
