package usecase

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
)

type Match struct {
	repo *repository.MatchRepository
}

func NewMatch(repo *repository.MatchRepository) *Match {
	return &Match{
		repo: repo,
	}
}

func (app *Match) FindMatch(playerID string, teamType domain.TeamType) (*domain.Match, bool, bool) {
	player := &domain.Player{ID: playerID}
	waitings := app.repo.WaitingMatches(teamType)

	var match domain.Match
	isTeam1 := true

	team := domain.NewTeam(teamType, player)
	if len(waitings) > 0 {
		match = *waitings[0]
		match.Join(&team)
		isTeam1 = false
	} else {
		match = domain.NewMatch(&team)
	}

	match.Start()
	app.repo.Save(&match)

	return &match, isTeam1, match.IsMatched()
}

func (app *Match) JoinMatch(matchID string, playerID string) *domain.Match {
	player := domain.Player{ID: playerID}
	match := app.repo.Find(matchID)
	if match == nil {
		return nil
	}

	match.MarkReady(player.ID)
	app.repo.Save(match)

	return match
}
