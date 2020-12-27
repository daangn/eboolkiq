package filedb

import (
	"context"
	"errors"
	"os"
	"testing"

	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
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

func TestFileDB_GetQueue(t *testing.T) {
	db := &FileDB{
		baseDir: "test",
		dbmap:   map[string]*bbolt.DB{},
	}
	tests := []struct {
		name   string
		search string
		queue  *pb.Queue
		err    error
	}{
		{
			name:   "normal case",
			search: "test",
			queue:  &pb.Queue{Name: "test"},
			err:    nil,
		}, {
			name:   "when queue not exists",
			search: "notfound",
			queue:  nil,
			err:    eboolkiq.ErrQueueNotFound,
		},
	}

	for _, test := range tests {
		queue, err := db.GetQueue(context.Background(), test.search)

		if !errors.Is(err, test.err) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n", test.name, test.err, err)
		}

		if err != nil {
			continue
		}

		if !proto.Equal(test.queue, queue) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n", test.name, test.queue, queue)
		}
	}
}

func TestFileDB_CreateQueue(t *testing.T) {
	if err := os.MkdirAll("test/temp", 0775); err != nil {
		t.Fatalf("fail to prepare test: %+v\n", err)
		return
	}
	defer func() {
		if err := os.RemoveAll("test/temp"); err != nil {
			t.Fatalf("fail to cleanup test: %+v\n", err)
		}
	}()

	db := &FileDB{
		baseDir: "test/temp",
		dbmap:   map[string]*bbolt.DB{},
	}
	tests := []struct {
		name  string
		queue *pb.Queue
		err   error
	}{
		{
			name:  "normal case",
			queue: &pb.Queue{Name: "test"},
			err:   nil,
		}, {
			name:  "create exists queue",
			queue: &pb.Queue{Name: "test"},
			err:   eboolkiq.ErrQueueExists,
		}, {
			name:  "create anonymous queue",
			queue: &pb.Queue{Name: ""},
			err:   bbolt.ErrKeyRequired,
		},
	}

	for _, test := range tests {
		queue, err := db.CreateQueue(context.Background(), test.queue)

		if !errors.Is(err, test.err) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n", test.name, test.err, err)
		}

		if err != nil {
			continue
		}

		if !proto.Equal(test.queue, queue) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n", test.name, test.queue, queue)
		}
	}
}

func TestFileDB_UpdateQueue(t *testing.T) {
	if err := os.MkdirAll("test/temp", 0775); err != nil {
		t.Fatalf("fail to prepare test: %+v\n", err)
		return
	}
	defer func() {
		if err := os.RemoveAll("test/temp"); err != nil {
			t.Fatalf("fail to cleanup test: %+v\n", err)
		}
	}()

	db := &FileDB{
		baseDir: "test/temp",
		dbmap:   map[string]*bbolt.DB{},
	}

	if _, err := db.CreateQueue(context.Background(), &pb.Queue{Name: "test"}); err != nil {
		t.Fatalf("fail to prepare test: %+v\n", err)
	}

	tests := []struct {
		name  string
		queue *pb.Queue
		err   error
	}{
		{
			name:  "normal case",
			queue: &pb.Queue{Name: "test", MaxRetry: 3},
			err:   nil,
		}, {
			name:  "try update anonymous queue",
			queue: &pb.Queue{Name: "anonymous"},
			err:   eboolkiq.ErrQueueNotFound,
		},
	}

	for _, test := range tests {
		queue, err := db.UpdateQueue(context.Background(), test.queue)

		if !errors.Is(err, test.err) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n", test.name, test.err, err)
		}

		if err != nil {
			continue
		}

		if !proto.Equal(queue, test.queue) {
			t.Errorf("test failed\n"+
				"case:     %+v\n"+
				"expected: %+v\n"+
				"actual:   %+v\n", test.name, test.queue, queue)
		}
	}
}
