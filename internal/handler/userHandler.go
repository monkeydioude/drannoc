package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/body"
	"github.com/monkeydioude/drannoc/internal/config"
	"github.com/monkeydioude/drannoc/internal/entity"
	repo "github.com/monkeydioude/drannoc/internal/repository"
	service "github.com/monkeydioude/drannoc/internal/routine"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// UserCreate handles user creation form
// POST /user
func UserCreate(c *gin.Context) {
	loginData, err := body.NewLoginData(c.Request.Body)
	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	token, err := service.UserCreate(
		loginData,
		repo.NewUser(),
		repo.NewAuthToken(),
	)

	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}
	if token == nil {
		res.ServiceUnavailable("Could not create user", "Could not create user")
		return
	}
	res.Ok(c, gin.H{
		"data": token,
	})
}

// UserLogin handles user login form
// POST /user/login
func UserLogin(c *gin.Context) {
	loginData, err := body.NewLoginData(c.Request.Body)
	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	user, err := entity.NewUser(loginData.Login, loginData.Password)
	// could not generate new user
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not generate new user", err.Error()))
		return
	}
	_, err = repo.NewUser().Load(user)

	// user does not exist
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not retrieve user", err.Error()))
		return
	}

	// create a new auth token on login
	token, err := service.CreateAuthTokenNow(repo.NewAuthToken(), user.ID)

	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not create token", err.Error()))
		return
	}
	maxAge := int(token.Expires - time.Now().Unix())
	// setting up the AuthToken Cookie
	c.SetCookie(config.AuthTokenLabel, token.GetToken(), maxAge, "/", "", false, false)
	c.SetCookie(config.ConsumerLabel, user.ID, maxAge, "/", "", false, false)

	res.Ok(c, gin.H{
		"data": token,
	})
}

// UserIndex retrieves user related data
// GET /user
func UserIndex(c *gin.Context) {
}
