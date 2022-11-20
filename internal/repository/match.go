package repository

import (
	"errors"
	"sync"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository/schema"
)

var ErrMatchNotFound = errors.New("match not found")

type Matches interface {
	Find(id string) (*domain.Match, error)
	ListAvaiable(excludeTeam domain.TeamType) []*domain.Match
	Save(*domain.Match) error
}

type SimpleMatchRepository struct {
	mutex sync.RWMutex
	items map[string]schema.Match
}

func NewSimpleMatchRepository() *SimpleMatchRepository {
	return &SimpleMatchRepository{
		items: make(map[string]schema.Match),
	}
}

func (repo *SimpleMatchRepository) Find(id string) (*domain.Match, error) {
	repo.mutex.RLock()
	data, ok := repo.items[id]
	repo.mutex.RUnlock()
	if ok == false {
		return nil, ErrMatchNotFound
	}

	return createMatchEntity(&data), nil
}

func (repo *SimpleMatchRepository) ListAvaiable(excludeTeam domain.TeamType) []*domain.Match {
	matches := []*domain.Match{}
	repo.mutex.RLock()
	for _, item := range repo.items {
		isCreated := item.State == domain.MatchCreated
		isOpponent := item.Team1Type != excludeTeam
		if isCreated && isOpponent {
			matches = append(matches, createMatchEntity(&item))
		}
	}
	repo.mutex.RUnlock()

	return matches
}

func (repo *SimpleMatchRepository) Save(match *domain.Match) error {
	repo.mutex.Lock()
	repo.items[match.ID] = schema.Match{
		ID:    match.ID,
		State: match.State(),

		Team1ID:    match.Team1().ID(),
		Team1Type:  match.Team1().Type,
		Team1Ready: match.Team1().IsReady,

		Team2ID:    match.Team2().ID(),
		Team2Type:  match.Team2().Type,
		Team2Ready: match.Team2().IsReady,
	}
	repo.mutex.Unlock()

	return nil
}

func createMatchEntity(data *schema.Match) *domain.Match {
	player1 := domain.NewPlayer(data.Team1ID)
	player2 := domain.NewPlayer(data.Team2ID)

	team1 := domain.NewTeam(data.Team1Type, &player1)
	team2 := domain.NewTeam(data.Team2Type, &player2)

	if data.Team1Ready {
		team1.ToReady()
	}

	if data.Team2Ready {
		team2.ToReady()
	}

	match := domain.NewMatchFromData(
		data.ID,
		data.State,
		&team1,
		&team2,
	)

	return &match
}
