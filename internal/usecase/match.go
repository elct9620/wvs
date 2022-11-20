package usecase

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
)

type Match struct {
	matches repository.Matches
}

func NewMatch(matches repository.Matches) *Match {
	return &Match{
		matches: matches,
	}
}

func (usecase *Match) FindMatch(playerID string, teamType domain.TeamType) (*domain.Match, bool, bool) {
	player := &domain.Player{ID: playerID}
	availableMatchs := usecase.matches.ListAvaiable(teamType)

	var match domain.Match
	isTeam1 := true

	team := domain.NewTeam(teamType, player)
	if len(availableMatchs) > 0 {
		match = *availableMatchs[0]
		match.Join(&team)
		isTeam1 = false
	} else {
		match = domain.NewMatch(&team)
	}

	match.Start()
	usecase.matches.Save(&match)

	return &match, isTeam1, match.IsMatched()
}

func (usecase *Match) JoinMatch(matchID string, playerID string) *domain.Match {
	player := domain.Player{ID: playerID}
	match, err := usecase.matches.Find(matchID)
	if err != nil {
		return nil
	}

	match.MarkReady(player.ID)
	usecase.matches.Save(match)

	return match
}
