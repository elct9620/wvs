package repository

import (
	"sync"

	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/internal/usecase"
)

type inMemoryMatchPlayerSchema struct {
	ID   string
	Team match.Team
}

type inMemoryMatchSchema struct {
	ID      string
	Players []inMemoryMatchPlayerSchema
}

var _ usecase.MatchRepository = &InMemoryMatchRepository{}

type InMemoryMatchRepository struct {
	mux     sync.RWMutex
	matches map[string]inMemoryMatchSchema
}

func NewInMemoryMatchRepository() *InMemoryMatchRepository {
	return &InMemoryMatchRepository{
		matches: make(map[string]inMemoryMatchSchema),
	}
}

func (r *InMemoryMatchRepository) Save(match *match.Match) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	players := make([]inMemoryMatchPlayerSchema, 0, len(match.Players()))
	for _, player := range match.Players() {
		players = append(players, inMemoryMatchPlayerSchema{
			ID:   player.Id(),
			Team: player.Team(),
		})
	}

	r.matches[match.Id()] = inMemoryMatchSchema{
		ID:      match.Id(),
		Players: players,
	}

	return nil
}
