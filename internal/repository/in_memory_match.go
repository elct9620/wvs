package repository

import (
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

func (r *InMemoryMatchRepository) Save(entity *match.Match) error {
	tnx := r.memdb.Txn(true)
	defer tnx.Abort()

	players := make([]db.MatchPlayer, 0, len(entity.Players()))
	for _, player := range entity.Players() {
		players = append(players, db.MatchPlayer{
			Id:   player.Id(),
			Team: int(player.Team()),
		})
	}

	match := &db.Match{
		Id:      entity.Id(),
		Players: players,
	}

	if err := tnx.Insert(db.TableMatch, match); err != nil {
		return err
	}

	tnx.Commit()
	return nil
}
