package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var databases map[string]*mongo.Database

// AuthDbName @TODO retrieve this from env var
var AuthDbName string = "drannoc"
var CoinsDbName string = "coins"

// Start instantiate MongoDB engine.
// Should be called at the very beg
func Start(addr string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))

	if err != nil {
		panic(err)
	}

	err = cl.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	databases = make(map[string]*mongo.Database)
	client = cl
	return client
}

// Database returns a mongo.Database by first
// looking into the singleton cached databasses list
func Database(db string) *mongo.Database {
	if _, ok := databases[db]; !ok {
		databases[db] = client.Database(db)
	}
	return databases[db]
}

// GetClient returns *mongo.Client singleton
func GetClient() *mongo.Client {
	return client
}
