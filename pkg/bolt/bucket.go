package bolt

// Bucket describes a bolt bucket. No need for interface ?
type Bucket struct {
	name []byte
	db   DatabaseTransaction
}

// Do executes a Transaction
func (b *Bucket) Do(t Transaction) ([]byte, error) {
	return t.Transaction(b)
}

// DB returns DB *struct
func (b *Bucket) DB() DatabaseTransaction {
	return b.db
}

// Name gives bucket's ~~head~~ name
func (b *Bucket) Name() []byte {
	return b.name
}

// Get is a helper for calling a view Transaction
func (b *Bucket) Get(key string) ([]byte, error) {
	return b.Do(Get(key))
}

// Put is a helper for calling an update Transaction
func (b *Bucket) Put(key, value string) ([]byte, error) {
	return b.Do(Put(key, value))
}

// Store handles putting Storable struct into a bucket
func (b *Bucket) Store(entity Entity) ([]byte, error) {
	return b.Put(entity.GetKey(), entity.String())
}

// NewBucket instantiate a Bucket struct
func NewBucket(name string, db DatabaseTransaction) *Bucket {
	return &Bucket{
		name: []byte(name),
		db:   db,
	}
}
