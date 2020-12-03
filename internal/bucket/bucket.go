package bucket

import "github.com/monkeydioude/drannoc/pkg/bolt"

// Auth is a helper, sort of a repository pattern
// responsible for knowing how to store Auth entities
func Auth(db bolt.DatabaseTransaction) *bolt.Bucket {
	if db == nil {
		db = bolt.NewDatabase("../../db/drannoc.db.bson")
	}
	return bolt.NewBucket("auth", db)
}

// AuthToken is a helper, sort of a repository pattern
// responsible for knowing how to store Auth entities
func AuthToken(db bolt.DatabaseTransaction) *bolt.Bucket {
	if db == nil {
		db = bolt.NewDatabase("../../db/drannoc.db.bson")
	}
	return bolt.NewBucket("authToken", db)
}

// Config holds all the static config of the system
func Config(db bolt.DatabaseTransaction) *bolt.Bucket {
	if db == nil {
		db = bolt.NewDatabase("../../db/drannoc.db.bson")
	}
	return bolt.NewBucket("config", db)
}

// User holds user's data
func User(db bolt.DatabaseTransaction) *bolt.Bucket {
	if db == nil {
		db = bolt.NewDatabase("../../db/drannoc.db.bson")
	}
	return bolt.NewBucket("user", db)
}
