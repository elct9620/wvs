package db

import "github.com/hashicorp/go-memdb"

const (
	TableBattle         = "battle"
	IndexBattleId       = "id"
	IndexBattleMatchId  = "matchId"
	IndexBattlePlayerId = "playerId"
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
		IndexBattleMatchId: {
			Name: IndexBattleMatchId,
			Indexer: &memdb.StringFieldIndex{
				Field: "MatchId",
			},
		},
	},
}

type BattleEvent struct {
	Id      string
	MatchId string
}
