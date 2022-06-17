package repository

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
)

type schema struct {
	ID    string
	State domain.MatchState

	Player1ID   string
	Player1Team domain.TeamType

	Player2ID   string
	Player2Team domain.TeamType
}

type MatchRepository struct {
	store *store.Store
}

func NewMatchRepository(store *store.Store) *MatchRepository {
	return &MatchRepository{
		store: store,
	}
}

func (repo *MatchRepository) Find(id string) *domain.Match {
	matches := repo.store.Table("matches")
	raw, err := matches.Find(id)
	if err != nil {
		return nil
	}
	data := raw.(schema)

	team1 := domain.NewTeam(data.Player1Team, &domain.Player{ID: data.Player1ID})
	team2 := domain.NewTeam(data.Player2Team, &domain.Player{ID: data.Player2ID})
	match := domain.NewMatchFromData(data.ID, data.State, &team1, &team2)
	return &match
}

func (repo *MatchRepository) Save(match domain.Match) error {
	matches := repo.store.Table("matches")

	return matches.Update(match.ID, schema{
		ID:          match.ID,
		State:       match.State(),
		Player1ID:   match.Player1().ID(),
		Player1Team: match.Player1().Type,
		Player2ID:   match.Player2().ID(),
		Player2Team: match.Player2().Type,
	})
}

func (repo *MatchRepository) WaitingMatches(excludeTeam domain.TeamType) []domain.Match {
	return []domain.Match{}
}
