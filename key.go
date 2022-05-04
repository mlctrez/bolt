package bolt

import (
	"go.etcd.io/bbolt"
)

type Key string

func (b Key) B() []byte {
	return []byte(b)
}

func (b Key) CreateBucket(tx *bbolt.Tx) error {
	_, err := tx.CreateBucketIfNotExists(b.B())
	return err
}

func (b Key) Bucket(tx *bbolt.Tx) *Bucket {
	if bBucket := tx.Bucket(b.B()); bBucket != nil {
		return &Bucket{bBucket}
	}
	return nil
}

type Keys []Key

func (keys Keys) CreateBucket(tx *bbolt.Tx) error {
	for _, key := range keys {
		if err := key.CreateBucket(tx); err != nil {
			return err
		}
	}
	return nil
}

type HasKey interface {
	Key() Key
}

type ValueProvider interface {
	HasKey
	Value() ([]byte, error)
}

type ValueReceiver interface {
	HasKey
	SetValue([]byte) error
}
