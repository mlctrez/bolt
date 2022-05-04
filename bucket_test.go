package bolt

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

var _ ValueProvider = (*testValue)(nil)
var _ ValueReceiver = (*testValue)(nil)

type testValue struct {
	k Key
	v []byte
}

func (v *testValue) Key() Key {
	return v.k
}
func (v *testValue) Value() ([]byte, error) {
	return v.v, nil
}
func (v *testValue) SetValue(bytes []byte) error {
	v.v = bytes
	return nil
}

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

	testVal := &testValue{k: key, v: valueBytes}
	req.Nil(bolt.Update(func(tx *bbolt.Tx) error {
		return bucket.Bucket(tx).Put(testVal)
	}))

	readVal := &testValue{k: key}
	req.Nil(bolt.View(func(tx *bbolt.Tx) error {
		req.Nil(bucket.Bucket(tx).Get(readVal))
		if readVal.v == nil {
			return fmt.Errorf("value not read")
		}
		return nil
	}))
	req.Equal(valueBytes, readVal.v)

}
