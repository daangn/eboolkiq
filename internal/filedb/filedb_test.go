package filedb

import (
	"errors"
	"os"
	"testing"

	"go.etcd.io/bbolt"
)

func TestFileDB_openDB(t *testing.T) {
	fileDB := &FileDB{
		dbmap: map[string]*bbolt.DB{},
	}

	tests := []struct {
		name string
		path string
		err  error
	}{
		{
			name: "normal case",
			path: "test/sample.db",
			err:  nil,
		}, {
			name: "reopen normal case",
			path: "test/sample.db",
			err:  nil,
		}, {
			name: "try open 0400 file",
			path: "test/cannot_open.db",
			err:  os.ErrPermission,
		},
	}

	for _, test := range tests {
		_, err := fileDB.openDB(test.path)

		if !errors.Is(err, test.err) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n"+
				"with:     %+v\n", test.name, test.err, err, test)
		}
	}
}
