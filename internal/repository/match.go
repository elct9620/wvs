package repository

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
)

type schema struct {
	ID    string
	State domain.MatchState

	Team1ID   string
	Team1Team domain.TeamType

	Team2ID   string
	Team2Team domain.TeamType
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

	team1 := domain.NewTeam(data.Team1Team, &domain.Player{ID: data.Team1ID})
	team2 := domain.NewTeam(data.Team2Team, &domain.Player{ID: data.Team2ID})
	match := domain.NewMatchFromData(data.ID, data.State, &team1, &team2)
	return &match
}

func (repo *MatchRepository) WaitingMatches(excludeTeam domain.TeamType) []*domain.Match {
	items := repo.store.Table("matches").Map(func(raw interface{}) interface{} {
		data := raw.(schema)

		team1 := domain.NewTeam(data.Team1Team, &domain.Player{ID: data.Team1ID})
		team2 := domain.NewTeam(data.Team2Team, &domain.Player{ID: data.Team2ID})
		return domain.NewMatchFromData(data.ID, data.State, &team1, &team2)
	})

	filtered := make([]*domain.Match, 0)
	for _, item := range items {
		match := item.(domain.Match)
		if match.State() != domain.MatchCreated || !match.Team1().IsValid() || match.Team2().IsValid() {
			continue
		}

		if match.Team1().Type == excludeTeam {
			continue
		}

		filtered = append(filtered, &match)
	}

	return filtered
}

func (repo *MatchRepository) Save(match *domain.Match) error {
	matches := repo.store.Table("matches")

	return matches.Update(match.ID, schema{
		ID:        match.ID,
		State:     match.State(),
		Team1ID:   match.Team1().ID(),
		Team1Team: match.Team1().Type,
		Team2ID:   match.Team2().ID(),
		Team2Team: match.Team2().Type,
	})
}
