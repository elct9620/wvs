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

	return &Database{MemDB: db}, nil
}
