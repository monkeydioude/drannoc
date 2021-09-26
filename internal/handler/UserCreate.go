package handler

import (
	"github.com/gin-gonic/gin"
	repo "github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
	"github.com/monkeydioude/drannoc/internal/service"
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
