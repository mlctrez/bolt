package bolt

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

func TestBolt_New(t *testing.T) {
	req := require.New(t)
	req.True(true)

	bolt, err := New(filepath.Join(t.TempDir(), "bolt.db"))
	req.Nil(err)
	req.NotNil(bolt)

	defer func() { req.Nil(bolt.Close()) }()
}

func TestBolt_CreateBuckets(t *testing.T) {
	req := require.New(t)
	req.True(true)

	bolt, err := New(filepath.Join(t.TempDir(), "bolt.db"))
	req.Nil(err)
	req.NotNil(bolt)
	defer func() { req.Nil(bolt.Close()) }()

	aBucket := Key("bucketOne")

	req.Nil(bolt.CreateBuckets(Keys{aBucket}))
	req.Nil(bolt.bBoltDB().View(func(tx *bbolt.Tx) error {
		if bucket := aBucket.Bucket(tx); bucket == nil {
			return fmt.Errorf("bucket %s was not created", aBucket)
		}
		return nil
	}))

}
