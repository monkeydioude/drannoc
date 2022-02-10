package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	in_gin "github.com/monkeydioude/drannoc/internal/gin"
	repo "github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
)

// UserIndex retrieves user related data
// GET /user
func UserIndex(c *gin.Context) {
	userID := in_gin.GetUserIDFromContext(c)

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
