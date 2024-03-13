package repository

import (
	"context"
	"fmt"

	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/hashicorp/go-memdb"
)

var _ usecase.MatchRepository = &InMemoryMatchRepository{}

type InMemoryMatchRepository struct {
	memdb *memdb.MemDB
}

func NewInMemoryMatchRepository(memdb *memdb.MemDB) *InMemoryMatchRepository {
	return &InMemoryMatchRepository{
		memdb: memdb,
	}
}

func (r *InMemoryMatchRepository) WaitingList(ctx context.Context) ([]*match.Match, error) {
	tnx := r.memdb.Txn(false)
	defer tnx.Abort()

	iter, err := tnx.Get(db.TableMatch, db.IndexMatchIsWaiting)
	if err != nil {
		return nil, err
	}

	matches := make([]*match.Match, 0)
	for {
		raw := iter.Next()
		if raw == nil {
			break
		}

		record, ok := raw.(*db.Match)
		if !ok {
			return nil, fmt.Errorf("unexpected type %T", record)
		}

		entity, err := newMatchFromDbRecord(record)
		if err != nil {
			continue
		}

		matches = append(matches, entity)
	}

	return matches, nil
}

func (r *InMemoryMatchRepository) Save(ctx context.Context, entity *match.Match) error {
	tnx := r.memdb.Txn(true)
	defer tnx.Abort()

	record, err := tnx.First(db.TableMatch, db.IndexMatchId, entity.Id())
	if err != nil {
		return err
	}

	var match *db.Match
	if record == nil {
		match = &db.Match{
			Id: entity.Id(),
		}
	} else {
		match, ok := record.(*db.Match)
		if !ok {
			return fmt.Errorf("unexpected type %T", match)
		}
	}

	players := make([]db.MatchPlayer, 0, len(entity.Players()))
	for _, player := range entity.Players() {
		players = append(players, db.MatchPlayer{
			Id:   player.Id(),
			Team: int(player.Team()),
		})
	}
	match.Players = players

	if err := tnx.Insert(db.TableMatch, match); err != nil {
		return err
	}

	tnx.Commit()
	return nil
}

func newMatchFromDbRecord(record *db.Match) (*match.Match, error) {
	entity := match.NewMatch(record.Id)
	for _, player := range record.Players {
		if err := entity.AddPlayer(player.Id, match.Team(player.Team)); err != nil {
			return nil, err
		}
	}

	return entity, nil
}
