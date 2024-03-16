package inmemory

import (
	"context"

	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/internal/usecase"
)

var _ usecase.MatchRepository = &MatchRepository{}

type MatchRepository struct {
	db *db.Database
}

func NewMatchRepository(db *db.Database) *MatchRepository {
	return &MatchRepository{
		db: db,
	}
}

func (r *MatchRepository) FindByPlayerID(ctx context.Context, playerId string) (*match.Match, error) {
	tnx := r.db.Txn(false)
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
	tnx := r.db.Txn(false)
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
	tnx := r.db.Txn(true)
	defer tnx.Abort()

	match := matchToRecord(entity)
	if err := tnx.Insert(db.TableMatch, match); err != nil {
		return err
	}

	tnx.Commit()
	return nil
}
