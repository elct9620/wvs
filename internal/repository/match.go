package repository

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
)

type MatchRepository struct {
	store *store.Table
}

func NewMatchRepository(store *store.Table) *MatchRepository {
	return &MatchRepository{
		store: store,
	}
}

func (repo *MatchRepository) WaitingMatches(excludeTeam domain.Team) []domain.Match {
	return []domain.Match{}
}
