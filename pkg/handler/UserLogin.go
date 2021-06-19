package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/entity"
	repo "github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"github.com/monkeydioude/drannoc/pkg/service"
)

// UserLogin handles user login form
// POST /user/login
func UserLogin(c *gin.Context) {
	user := &entity.User{}
	err := service.EntityFromRequestBody(c.Request.Body, user)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not generate new user", err.Error()))
		return
	}
	user.PasswordEncrypt()

	exist, err := repo.NewUser().LoadFromCredentials(user)

	// error retrieving
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not retrieve user", err.Error()))
		return
	}

	if exist == nil {
		res.Write(c, res.BadRequest("user does not exist"))
		return
	}

	// create a new auth token on login
	token := service.CreateAuthTokenNow(user.ID, c.GetInt("TokenLivesMaxAmount"))
	_, err = repo.NewAuthToken().Store(token)

	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not create token", err.Error()))
		return
	}

	repoPref := repo.NewUserPreferences()
	pref, _ := service.FetchUserPreferences(repoPref, user.ID)

	service.SetCookies(token, user.GetID(), c)

	res.Ok(c, gin.H{
		"data": res.UserIndex{
			User:        nil,
			Preferences: pref,
		},
	})
}
