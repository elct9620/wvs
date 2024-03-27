package db

import (
	"encoding/binary"
	"fmt"

	"github.com/hashicorp/go-memdb"
)

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
			Name:    IndexBattleAggregateId,
			Unique:  true,
			Indexer: &BattleVersionIndexer{},
		},
	},
}

type BattleEvent struct {
	Id          string
	AggregateId string
	Type        string
	Version     int
	CreatedAt   int64
}

var _ memdb.Indexer = &BattleVersionIndexer{}
var _ memdb.MultiIndexer = &BattleVersionIndexer{}

type BattleVersionIndexer struct{}

func (i *BattleVersionIndexer) FromArgs(args ...any) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	aggregateId, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("expected string argument, got %T", args[0])
	}

	return []byte(aggregateId), nil
}

func (i *BattleVersionIndexer) FromObject(obj any) (bool, [][]byte, error) {
	event, ok := obj.(*BattleEvent)
	if !ok {
		return false, nil, fmt.Errorf("expected *BattleEvent object, got %T", obj)
	}

	index := make([][]byte, 2)
	index[0] = []byte(event.AggregateId)
	index[1] = make([]byte, 8)
	binary.BigEndian.PutUint64(index[1], uint64(event.Version))

	return true, index, nil
}
