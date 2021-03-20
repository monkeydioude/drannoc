package config

import "os"

const (
	AuthTokenLabel string = "AT"
	// OriginDomain   string = os.Getenv("ORIGIN_DOMAIN")
	UserLoginRoute = "/user/login"
)

var OriginDomain string = os.Getenv("ORIGIN_DOMAIN")
var MongoDBAddr = os.Getenv("MONGODB_ADDR")

func init() {
	OriginDomain = os.Getenv("ORIGIN_DOMAIN")
	if OriginDomain == "" {
		panic("ORIGIN_DOMAIN env variable required")
	}
}
