package config

import "os"

const (
	// OriginDomain   string = os.Getenv("ORIGIN_DOMAIN")
	AuthTokenLabel      string = "auth-token"
	ConsumerLabel       string = "consumer"
	UserLoginRoute      string = "/user/login"
	TokenLivesMaxAmount int    = 5
)

var OriginDomain string = os.Getenv("ORIGIN_DOMAIN")
var MongoDBAddr = os.Getenv("MONGODB_ADDR")
var ServerPort = os.Getenv("SERVER_PORT")

func init() {
	OriginDomain = os.Getenv("ORIGIN_DOMAIN")
}
