package repository

import "github.com/hashicorp/go-memdb"

type TableSchemaFn func(schema *memdb.DBSchema)

func NewMemDB() (*memdb.MemDB, error) {
	schema := &memdb.DBSchema{
		Tables: make(map[string]*memdb.TableSchema),
	}

	tableSchemas := []TableSchemaFn{
		defineMemoryRoomTable,
		defineMemoryPlayerTable,
	}

	for _, fn := range tableSchemas {
		fn(schema)
	}

	return memdb.NewMemDB(schema)
}

func defineMemoryRoomTable(schema *memdb.DBSchema) {
	schema.Tables[RoomTableName] = &memdb.TableSchema{
		Name: RoomTableName,
		Indexes: map[string]*memdb.IndexSchema{
			"id": {
				Name:    "id",
				Unique:  true,
				Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
			},
			"state": {
				Name:    "state",
				Indexer: &memdb.IntFieldIndex{Field: "State"},
			},
		},
	}
}

func defineMemoryPlayerTable(schema *memdb.DBSchema) {
	schema.Tables[PlayerTableName] = &memdb.TableSchema{
		Name: PlayerTableName,
		Indexes: map[string]*memdb.IndexSchema{
			"id": {
				Name:    "id",
				Unique:  true,
				Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
			},
		},
	}
}
