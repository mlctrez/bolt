package bolt

import (
	"time"

	"go.etcd.io/bbolt"
)

type Bolt struct {
	db *bbolt.DB
}

func New(path string) (*Bolt, error) {
	boltOptions := &bbolt.Options{Timeout: 5 * time.Second}
	if db, err := bbolt.Open(path, 0644, boltOptions); err != nil {
		return nil, err
	} else {
		return &Bolt{db: db}, nil
	}
}

func (s *Bolt) bBoltDB() *bbolt.DB {
	return s.db
}

func (s *Bolt) Close() error {
	return s.db.Close()
}

func (s *Bolt) CreateBuckets(buckets Keys) error {
	return s.db.Update(buckets.CreateBucket)
}

func (s *Bolt) Update(update func(tx *bbolt.Tx) error) error {
	return s.db.Update(update)
}

func (s *Bolt) View(view func(tx *bbolt.Tx) error) error {
	return s.db.View(view)
}

func (s *Bolt) Put(bucket Key, provider ValueProvider) error {
	return s.Update(func(tx *bbolt.Tx) error {
		return bucket.Bucket(tx).Put(provider)
	})
}

func (s *Bolt) Get(bucket Key, receiver ValueReceiver) error {
	return s.View(func(tx *bbolt.Tx) error {
		return bucket.Bucket(tx).Get(receiver)
	})
}
