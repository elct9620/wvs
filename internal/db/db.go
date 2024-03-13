package db

import (
	"github.com/google/wire"
	"github.com/hashicorp/go-memdb"
)

var DefaultSet = wire.NewSet(
	ProvideDatabaseSchema,
	memdb.NewMemDB,
)

const (
	TableMatch = "match"
)

func ProvideDatabaseSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			TableMatch: MatchTableSchema,
		},
	}
}
