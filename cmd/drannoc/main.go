package main

import (
	"fmt"

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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, auth-token, consumer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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

	db.Start(config.MongoDBAddr)

	// r.POST("/auth/identify", handler.Authenticate)
	// User creation
	r.POST("/user", handler.UserCreate)
	r.POST(config.UserLoginRoute, handler.UserLogin)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/user", handler.UserIndex)
		authorized.DELETE("/auth", handler.AuthDelete)
	}

	r.Run(fmt.Sprintf(":%s", config.ServerPort))
}