package inmemory

import (
	"context"

	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/hashicorp/go-memdb"
)

var _ usecase.MatchRepository = &MatchRepository{}

type MatchRepository struct {
	memdb *memdb.MemDB
}

func NewMatchRepository(memdb *memdb.MemDB) *MatchRepository {
	return &MatchRepository{
		memdb: memdb,
	}
}

func (r *MatchRepository) FindByPlayerID(ctx context.Context, playerId string) (*match.Match, error) {
	tnx := r.memdb.Txn(false)
	defer tnx.Abort()

	raw, err := tnx.First(db.TableMatch, db.IndexMatchPlayerId, playerId)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	return recordToMatch(raw.(*db.Match))
}

func (r *MatchRepository) Waiting(ctx context.Context) ([]*match.Match, error) {
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

		entity, err := recordToMatch(raw.(*db.Match))
		if err != nil {
			continue
		}

		matches = append(matches, entity)
	}

	return matches, nil
}

func (r *MatchRepository) Save(ctx context.Context, entity *match.Match) error {
	tnx := r.memdb.Txn(true)
	defer tnx.Abort()

	match := matchToRecord(entity)
	if err := tnx.Insert(db.TableMatch, match); err != nil {
		return err
	}

	tnx.Commit()
	return nil
}
