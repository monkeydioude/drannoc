package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	in_gin "github.com/monkeydioude/drannoc/internal/gin"
	repo "github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
)

func UserPreferencesGet(c *gin.Context) {
	userID := in_gin.GetUserIDFromContext(c)

	if userID == "" {
		res.Write(c, res.BadRequest("could not find userID"))
		return
	}

	pref := &entity.UserPreferences{}
	_, err := repo.NewUserPreferences().Load(pref, userID)

	if err != nil {
		res.Write(c, res.BadRequest(fmt.Sprintf("cound not find preferences for this userID %s", userID)))
		return
	}
	res.Ok(c, gin.H{
		"data": pref,
	})
}
