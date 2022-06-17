package repository

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
)

type MatchRepository struct {
	store *store.Store
}

func NewMatchRepository(store *store.Store) *MatchRepository {
	return &MatchRepository{
		store: store,
	}
}

func (repo *MatchRepository) WaitingMatches(excludeTeam domain.TeamType) []domain.Match {
	return []domain.Match{}
}
