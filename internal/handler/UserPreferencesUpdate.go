package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	repo "github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
	"github.com/monkeydioude/drannoc/internal/service"
)

// UserPreferencesUpdate modifies a user's data
// in the Preferences object of a User
// PUT /user/preferences
func UserPreferencesUpdate(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.BadRequest("could not find userID"))
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
