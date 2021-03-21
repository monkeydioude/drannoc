package config

import "os"

const (
	AuthTokenLabel string = "auth-token"
	ConsumerLabel  string = "consumer"
	// OriginDomain   string = os.Getenv("ORIGIN_DOMAIN")
	UserLoginRoute = "/user/login"
)

var OriginDomain string = os.Getenv("ORIGIN_DOMAIN")
var MongoDBAddr = os.Getenv("MONGODB_ADDR")
var ServerPort = os.Getenv("SERVER_PORT")

func init() {
	OriginDomain = os.Getenv("ORIGIN_DOMAIN")
	if OriginDomain == "" {
		panic("ORIGIN_DOMAIN env variable required")
	}
}
