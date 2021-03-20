package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/config"
	"github.com/monkeydioude/drannoc/internal/db"
	"github.com/monkeydioude/drannoc/internal/handler"
	"github.com/monkeydioude/drannoc/internal/middleware"
)

// CORSMiddleware handles CORS rights and OPTIONS requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.OriginDomain)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// @todo dinamycally add allowed headers ?
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, AT")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	// r.POST("/auth/identify", handler.Authenticate)
	// User creation
	r.POST("/user", handler.UserCreate)
	r.POST(config.UserLoginRoute, handler.UserLogin)

	db.Start(config.MongoDBAddr)
	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/profile", handler.ProfileIndex)
		// authorized.GET("/test1", handler.TestHandler1)
	}

	// authorized.GET("/login", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "y a pas d'heure pour se faire plaiz",
	// 	})
	// })
	r.Run(":8081")
}
