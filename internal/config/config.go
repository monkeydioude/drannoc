package config

import (
	"os"
)

var MongoDBAddr = os.Getenv("MONGODB_ADDR")
var ServerPort = os.Getenv("SERVER_PORT")

const UserLoginRoute string = "/user/login"
