package bolt

const (
	// X Execution
	X int = 0100
	// W Write
	W = 0200
	// R Read
	R = 0400
	// RW Read and Write
	RW = R | W
	// RWX Read, Write and Execution
	RWX = R | W | X
)

// DatabaseTransaction defines a full transaction, from connexion to
// database operation
type DatabaseTransaction interface {
	Path() string
	Bucket(string) *Bucket
}

// Database contains data relatives to DB an
type Database struct {
	path string
}

// Path returns path of database (the bson file)
func (db *Database) Path() string {
	return db.path
}

// Bucket instantiante a new *Bucket struct, for chaining purpose
func (db *Database) Bucket(bucket string) *Bucket {
	return NewBucket(bucket, db)
}

// NewDatabase creates a new Database *struct
func NewDatabase(path string) *Database {
	return &Database{path}
}
