package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/handler"
)

func main() {
	r := gin.Default()

	// Such as login action
	r.POST("/auth", handler.Authenticate)
	// User creation
	r.POST("/user", handler.CreateUser)
	r.POST("/user/login", handler.UserLogin)

	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"wesh": "alors",
	// }))

	// authorized.GET("/login", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "y a pas d'heure pour se faire plaiz",
	// 	})
	// })
	r.Run()
}
