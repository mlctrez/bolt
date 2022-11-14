package bolt

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

var _ ValueProvider = (*errValueProvider)(nil)

type errValueProvider struct {
}

func (e *errValueProvider) Key() Key               { return "key" }
func (e *errValueProvider) Value() ([]byte, error) { return nil, fmt.Errorf("errValueProvider") }

func TestBucket_Put(t *testing.T) {
	req := require.New(t)
	req.True(true)

	bolt, err := New(filepath.Join(t.TempDir(), "bolt.db"))
	req.Nil(err)
	req.NotNil(bolt)
	defer func() { req.Nil(bolt.Close()) }()

	bucket := Key("bucketOne")
	req.Nil(bolt.CreateBuckets(Keys{bucket}))

	key := Key("abcd")
	valueBytes := []byte{0, 1, 2, 3}

	testVal := &Value{K: key, V: valueBytes}
	req.Nil(bolt.Update(func(tx *bbolt.Tx) error {
		return bucket.Bucket(tx).Put(testVal)
	}))

	readVal := &Value{K: key}
	req.Nil(bolt.View(func(tx *bbolt.Tx) error {
		req.Nil(bucket.Bucket(tx).Get(readVal))
		if readVal.V == nil {
			return fmt.Errorf("value not read")
		}
		return nil
	}))
	req.Equal(valueBytes, readVal.V)

	err = bolt.Put(bucket, &errValueProvider{})
	req.NotNil(err)

}

func TestBucket_Delete(t *testing.T) {

	req := require.New(t)
	req.True(true)

	bolt, err := New(filepath.Join(t.TempDir(), "bolt.db"))
	req.Nil(err)
	req.NotNil(bolt)
	defer func() { req.Nil(bolt.Close()) }()

	bucket := Key("bucketOne")
	req.Nil(bolt.CreateBuckets(Keys{bucket}))

	key := Key("abcd")
	valueBytes := []byte{0, 1, 2, 3}

	testVal := &Value{K: key, V: valueBytes}
	req.Nil(bolt.Put(bucket, testVal))

	req.Nil(bolt.Get(bucket, testVal))
	req.Nil(bolt.Delete(bucket, key))

	req.ErrorIs(ErrValueNotFound, bolt.Get(bucket, testVal))

}
