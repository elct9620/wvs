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
			TableMatch:  MatchTableSchema,
			TableBattle: BattleTableSchema,
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

func (d *Database) Txn(write bool) *memdb.Txn {
	txn := d.MemDB.Txn(write)
	txn.TrackChanges()
	txn.Defer(func() {
		changes := txn.Changes()
		if len(changes) > 0 {
			for _, change := range changes {
				go d.publish(&change)
			}
		}
	})

	return txn
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
