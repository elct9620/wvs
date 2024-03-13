package db

import (
	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/hashicorp/go-memdb"
)

const (
	IndexMatchId        = "id"
	IndexMatchIsWaiting = "isWaiting"
)

var MatchTableSchema = &memdb.TableSchema{
	Name: TableMatch,
	Indexes: map[string]*memdb.IndexSchema{
		IndexMatchId: {
			Name:   IndexMatchId,
			Unique: true,
			Indexer: &memdb.StringFieldIndex{
				Field: "Id",
			},
		},
		IndexMatchIsWaiting: {
			Name:   IndexMatchIsWaiting,
			Unique: false,
			Indexer: &memdb.ConditionalIndex{
				Conditional: MatchWaitingIndexFunc,
			},
		},
	},
}

type MatchPlayer struct {
	Id   string
	Team int
}

type Match struct {
	Id      string
	Players []MatchPlayer
}

func MatchWaitingIndexFunc(obj any) (bool, error) {
	record, ok := obj.(*Match)
	if !ok {
		return false, nil
	}

	return len(record.Players) < match.MaxPlayers, nil
}
