package repository

import "github.com/hashicorp/go-memdb"

func NewMemDB() (*memdb.MemDB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			RoomTableName: {
				Name: RoomTableName,
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	return memdb.NewMemDB(schema)
}