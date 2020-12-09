package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/handler"
	"github.com/monkeydioude/drannoc/internal/middleware"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Such as login action
	r.POST("/auth", handler.Authenticate)
	// User creation
	r.POST("/user", handler.UserCreate)
	r.POST("/user/login", handler.UserLogin)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/test1", handler.TestHandler1)
	}

	// authorized.GET("/login", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "y a pas d'heure pour se faire plaiz",
	// 	})
	// })
	r.Run()
}
