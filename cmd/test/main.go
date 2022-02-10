package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/config"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/handler"
	"github.com/monkeydioude/drannoc/pkg/db"
)

var CoinInfos []*entity.CoinInfo

func BuildConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("OriginDomain", os.Getenv("ORIGIN_DOMAIN"))
		c.Set("AuthTokenLabel", "auth-token")
		c.Set("ConsumerLabel", "consumer")
		c.Set("UserLoginRoute", "/user/login")
		c.Set("TokenLivesMaxAmount", 9999)
		c.Next()
	}
}

// CORSMiddleware handles CORS rights and OPTIONS requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetString("OriginDomain"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// @todo dinamycally add allowed headers ?
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With, auth-token, consumer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(BuildConfig())
	r.Use(gin.Recovery())

	db.Start(config.MongoDBAddr)

	// r.POST("/auth/identify", handler.Authenticate)
	// User creation
	r.POST("/user", handler.UserCreate)
	r.POST(config.UserLoginRoute, handler.UserLogin)

	r.GET("/coins/info", handler.CoinsInfo)
	r.GET("/coin/:coin_id", handler.CoinGet)
	r.GET("/coins", handler.CoinsGet)
	r.PUT("/user/preferences", handler.UserPreferencesUpdate)
	r.GET("/user/preferences", handler.UserPreferencesGet)
	r.GET("/user", handler.UserIndex)
	r.DELETE("/auth", handler.AuthDelete)
	r.PUT("/trade/:trade_id", handler.TradeEdit)
	r.DELETE("/trade/:trade_id", handler.TradeDelete)
	r.POST("/trade", handler.TradeAdd)
	r.GET("/trades", handler.TradesGet)

	r.Run(fmt.Sprintf(":%s", config.ServerPort))
}
