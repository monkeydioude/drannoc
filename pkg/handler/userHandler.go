package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/entity"
	repo "github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"github.com/monkeydioude/drannoc/pkg/service"
)

// UserCreate handles user creation form
// POST /user
func UserCreate(c *gin.Context) {
	err := service.UserCreate(
		c.Request.Body,
		repo.NewUser(),
		repo.NewAuthToken(),
	)

	if err != nil {
		res.Write(c, res.BadRequest(err.Error()))
		return
	}

	res.Ok(c, gin.H{})
}

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

	service.SetCookies(token, user.GetID(), c)

	res.Ok(c, gin.H{
		"data": token,
	})
}

// UserIndex retrieves user related data
// GET /user
func UserIndex(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.ServiceUnavailable("could not find userID", "no consumer in header"))
		return
	}

	user, err := repo.UserID(userID).Load()
	if err != nil {
		res.Write(c, res.ServiceUnavailable(err.Error(), err.Error()))
		return
	}

	prefs := entity.NewUserPreferences()
	_, err = repo.NewUserPreferences().Load(prefs, userID)
	if err != nil {
		res.Write(c, res.ServiceUnavailable(err.Error(), err.Error()))
		return
	}
	// obfuscate ID
	user.ID = ""
	prefs.ID = ""

	res.Ok(c, gin.H{
		"data": res.UserIndex{
			User:        user,
			Preferences: prefs,
		},
	})
}

// UserPreferencesUpdate modifies a user's data
// in the Preferences object of a User
// PUT /user/preferences
func UserPreferencesUpdate(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.ServiceUnavailable("could not find userID", "no consumer in header"))
		return
	}

	pref := &entity.UserPreferences{}
	err := service.EntityFromRequestBody(c.Request.Body, pref)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not generate new user", err.Error()))
		return
	}

	if userID != pref.ID {
		res.Write(c, res.Unauthorized("can not update user"))
		return
	}

	err = repo.NewUserPreferences().Save(pref)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not save user preferences", err.Error()))
		return
	}

	res.Ok(c, gin.H{})
}
