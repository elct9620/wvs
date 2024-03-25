package db

import "github.com/hashicorp/go-memdb"

const (
	TableBattle            = "battle"
	IndexBattleId          = "id"
	IndexBattleAggregateId = "aggregateId"
	IndexBattlePlayerId    = "playerId"
)

var BattleTableSchema = &memdb.TableSchema{
	Name: TableBattle,
	Indexes: map[string]*memdb.IndexSchema{
		IndexBattleId: {
			Name:   IndexBattleId,
			Unique: true,
			Indexer: &memdb.StringFieldIndex{
				Field: "Id",
			},
		},
		IndexBattleAggregateId: {
			Name: IndexBattleAggregateId,
			Indexer: &memdb.StringFieldIndex{
				Field: "AggregateId",
			},
		},
	},
}

type BattleEvent struct {
	Id          string
	AggregateId string
	Type        string
	CreatedAt   int64
}
