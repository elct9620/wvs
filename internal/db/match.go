package db

import (
	"fmt"

	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/hashicorp/go-memdb"
)

const (
	IndexMatchId        = "id"
	IndexMatchPlayerId  = "playerId"
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
		IndexMatchPlayerId: {
			Name:    IndexMatchPlayerId,
			Unique:  false,
			Indexer: &MatchPlayerIdIndexer{},
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

var _ memdb.Indexer = &MatchPlayerIdIndexer{}
var _ memdb.MultiIndexer = &MatchPlayerIdIndexer{}

type MatchPlayerIdIndexer struct{}

func (i *MatchPlayerIdIndexer) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	playerId, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("expected string, got %T", args[0])
	}

	return []byte(playerId), nil
}

func (i *MatchPlayerIdIndexer) FromObject(obj any) (bool, [][]byte, error) {
	record, ok := obj.(*Match)
	if !ok {
		return false, nil, fmt.Errorf("unexpected type %T", obj)
	}

	ids := make([][]byte, 0, len(record.Players))
	for _, player := range record.Players {
		ids = append(ids, []byte(player.Id))
	}

	return true, ids, nil
}
