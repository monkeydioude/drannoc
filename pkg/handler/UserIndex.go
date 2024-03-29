package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/entity"
	repo "github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

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
