package bolt

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

// Transaction defines a bolt transaction (update for ex)
type Transaction interface {
	Key() []byte
	Value() []byte
	Transaction(*Bucket) ([]byte, error)
}

// BaseTransaction is the most basic structure for a Transaction
type BaseTransaction struct {
	key   []byte
	value []byte
}

// Key implements Transaction interface
func (t *BaseTransaction) Key() []byte {
	return t.key
}

// Value implements Transaction interface
func (t *BaseTransaction) Value() []byte {
	return t.value
}

// PutTransaction wraps a boltDB update transaction
type PutTransaction struct {
	BaseTransaction
}

// Put creates a new *PutTransaction struct
func Put(key, value string) *PutTransaction {
	return &PutTransaction{
		BaseTransaction{
			key:   []byte(key),
			value: []byte(value),
		},
	}
}

// Transaction implements Transaction interface
// This transaction is meant to update a key inside a bucket and
// does not return any []byte value
func (t *PutTransaction) Transaction(bucket *Bucket) ([]byte, error) {
	boltdb, err := bolt.Open(bucket.DB().Path(), RW, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return nil, err
	}
	defer boltdb.Close()

	return nil, boltdb.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket.Name())
		if err != nil {
			return err
		}
		return b.Put(t.Key(), t.Value())
	})
}

// GetTransaction wraps a boltDB view transaction
type GetTransaction struct {
	BaseTransaction
}

// Transaction here describes a view call of a key inside a bucket
// and returns its content
func (t *GetTransaction) Transaction(bucket *Bucket) ([]byte, error) {
	var value []byte

	boltdb, err := bolt.Open(bucket.DB().Path(), RW, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return nil, err
	}
	defer boltdb.Close()

	err = boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket.Name())
		if b == nil {
			return fmt.Errorf("Bucket %s not found", string(bucket.Name()))
		}
		value = b.Get(t.Key())
		return nil
	})

	return value, err
}

// Get generates a GetTransaction *struct
func Get(key string) *GetTransaction {
	return &GetTransaction{
		BaseTransaction{
			key: []byte(key),
		},
	}
}
