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

	return dbRecordToMatch(raw.(*db.Match))
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

		entity, err := dbRecordToMatch(raw.(*db.Match))
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

	match := matchToDbRecord(entity)
	if err := tnx.Insert(db.TableMatch, match); err != nil {
		return err
	}

	tnx.Commit()
	return nil
}

func dbRecordToMatch(record *db.Match) (*match.Match, error) {
	entity := match.NewMatch(record.Id)
	for _, player := range record.Players {
		if err := entity.AddPlayer(player.Id, match.Team(player.Team)); err != nil {
			return nil, err
		}
	}

	return entity, nil
}

func matchToDbRecord(entity *match.Match) *db.Match {
	players := make([]db.MatchPlayer, 0, len(entity.Players()))
	for _, player := range entity.Players() {
		players = append(players, db.MatchPlayer{
			Id:   player.Id(),
			Team: int(player.Team()),
		})
	}

	return &db.Match{
		Id:      entity.Id(),
		Players: players,
	}
}
