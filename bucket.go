package bolt

import (
	"errors"

	"go.etcd.io/bbolt"
)

type Bucket struct {
	bucket *bbolt.Bucket
}

func (b *Bucket) bBucket() *bbolt.Bucket {
	return b.bucket
}

func (b *Bucket) Put(p ValueProvider) error {
	return b.bucket.Put(p.Key().B(), p.Value())
}

func (b *Bucket) Get(r ValueReceiver) error {
	data := b.bucket.Get(r.Key().B())
	if data == nil {
		return ErrValueNotFound
	}
	r.SetValue(data)
	return nil
}

var ErrValueNotFound = errors.New("value not found")
