package db

import (
	"github.com/google/wire"
	"github.com/hashicorp/go-memdb"
)

var DefaultSet = wire.NewSet(
	NewDatabase,
)

const (
	TableMatch = "match"
)

type Database struct {
	*memdb.MemDB
	watchers map[*Watcher]struct{}
}

func NewDatabase() (*Database, error) {
	db, err := memdb.NewMemDB(&memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			TableMatch: MatchTableSchema,
		},
	})

	if err != nil {
		return nil, err
	}

	return &Database{
		MemDB:    db,
		watchers: make(map[*Watcher]struct{}),
	}, nil
}

func (d *Database) Tnx(write bool) *memdb.Txn {
	tnx := d.MemDB.Txn(write)
	tnx.TrackChanges()
	tnx.Defer(func() {
		changes := tnx.Changes()
		if len(changes) > 0 {
			for _, change := range changes {
				d.publish(&change)
			}
		}
	})

	return tnx
}

func (d *Database) publish(change *memdb.Change) {
	for watcher := range d.watchers {
		if watcher.Closed() {
			delete(d.watchers, watcher)
			continue
		}

		watcher.ch <- change
	}
}

func (d *Database) Watch() *Watcher {
	watcher := NewWatcher()
	d.watchers[watcher] = struct{}{}

	return watcher
}
