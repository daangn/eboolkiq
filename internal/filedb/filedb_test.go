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
		db, err := fileDB.openDB(test.path)

		if !errors.Is(err, test.err) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n"+
				"with:     %+v\n", test.name, test.err, err, test)
		}

		if err != nil {
			continue
		}

		if err := db.Close(); err != nil {
			t.Errorf("unexpected error occured: %+v\n", err)
		}
	}
}

func TestFileDB_Close(t *testing.T) {
	db := &FileDB{
		dbmap: map[string]*bbolt.DB{},
	}

	if _, err := db.openDB("test/sample.db"); err != nil {
		t.Errorf("fail to prepare: %+v\n", err)
	}

	if err := db.Close(); err != nil {
		t.Errorf("test failed\n"+
			"expected: %+v\n"+
			"actual: %+v\n", nil, err)
	}
}

func TestFileDB_dbPath(t *testing.T) {
	tests := []struct {
		name string
		db   *FileDB
		path string
	}{
		{
			name: "normal case",
			db:   &FileDB{baseDir: ""},
			path: dbFile,
		}, {
			name: "contains relative path on baseDir",
			db:   &FileDB{baseDir: "test/../"},
			path: dbFile,
		}, {
			name: "root path",
			db:   &FileDB{baseDir: "/"},
			path: "/" + dbFile,
		},
	}

	for _, test := range tests {
		path := test.db.dbPath()

		if test.path != path {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n"+
				"with:     %+v\n", test.name, test.path, path, test)
		}
	}
}
