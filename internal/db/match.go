package db

import (
	"github.com/hashicorp/go-memdb"
)

const (
	IndexMatchId       = "id"
	IndexMatchPlayerId = "playerId"
)

var MatchTableSchema = &memdb.TableSchema{
	Name: "match",
	Indexes: map[string]*memdb.IndexSchema{
		IndexMatchId: {
			Name:   IndexMatchId,
			Unique: true,
			Indexer: &memdb.StringFieldIndex{
				Field: "Id",
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
