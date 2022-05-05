//go:build ignore

package main

import (
	"fmt"

	"github.com/mlctrez/bolt"
	"go.etcd.io/bbolt"
)

func main() {

	// error handling omitted for example
	db, _ := bolt.New("temp/bolt.db")
	defer func() { _ = db.Close() }()

	bucketOne := bolt.Key("BucketOne")
	_ = db.CreateBuckets(bolt.Keys{bucketOne})

	// single updates
	_ = db.Put(bucketOne, &stringStorage{key: "keyOne", value: "valueOne"})

	// multiple updates with transaction rollback
	_ = db.Update(func(tx *bbolt.Tx) error {
		bucketTx := bucketOne.Bucket(tx)
		for i := 0; i < 5; i++ {
			data := fmt.Sprintf("data_%d", i)
			err := bucketTx.Put(&stringStorage{key: bolt.Key(data), value: data})
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		return nil
	})

	// simple read
	simpleRead := &stringStorage{key: "keyOne"}
	_ = db.Get(bucketOne, simpleRead)
	fmt.Println(simpleRead.value)

	badLookup := &stringStorage{key: "doesnotexist"}
	if err := db.Get(bucketOne, badLookup); err == bolt.ErrValueNotFound {
		fmt.Println(err)
	}
}

var _ bolt.ValueProvider = (*stringStorage)(nil)
var _ bolt.ValueReceiver = (*stringStorage)(nil)

type stringStorage struct {
	key   bolt.Key
	value string
}

func (s *stringStorage) Value() ([]byte, error) {
	return []byte(s.value), nil
}

func (s *stringStorage) Key() bolt.Key {
	return s.key
}

func (s *stringStorage) SetValue(bytes []byte) error {
	s.value = string(bytes)
	return nil
}
