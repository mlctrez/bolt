package bolt

import (
	"fmt"
	"os"
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

func TestBolt_New_BadOpen(t *testing.T) {
	req := require.New(t)
	req.True(true)

	tempDir := t.TempDir()
	req.Nil(os.Chmod(tempDir, 0444))
	bolt, err := New(filepath.Join(tempDir, "bolt.db"))
	req.Nil(bolt)
	req.NotNil(err)

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
