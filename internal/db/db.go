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
	changes chan memdb.Change
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
		MemDB:   db,
		changes: make(chan memdb.Change),
	}, nil
}

func (d *Database) Tnx(write bool) *memdb.Txn {
	tnx := d.MemDB.Txn(write)
	tnx.TrackChanges()
	tnx.Defer(func() {
		changes := tnx.Changes()
		if len(changes) > 0 {
			for _, change := range changes {
				d.changes <- change
			}
		}
	})

	return tnx
}
