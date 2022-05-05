package bolt

import (
	"errors"

	"go.etcd.io/bbolt"
)

var ErrValueNotFound = errors.New("value not found")

type Bucket struct {
	bucket *bbolt.Bucket
}

func (b *Bucket) bBucket() *bbolt.Bucket {
	return b.bucket
}

func (b *Bucket) Delete(key Key) error {
	return b.bucket.Delete(key.B())
}

func (b *Bucket) Put(p ValueProvider) error {
	value, err := p.Value()
	if err != nil {
		return err
	}
	return b.bucket.Put(p.Key().B(), value)
}

func (b *Bucket) Get(r ValueReceiver) error {
	data := b.bucket.Get(r.Key().B())
	if data == nil {
		return ErrValueNotFound
	}
	return r.SetValue(data)
}
