package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
)

type MatchApplication struct {
	engine    *engine.Engine
	repo      *repository.MatchRepository
	broadcast *service.BroadcastService
}

func NewMatchApplication(engine *engine.Engine, repo *repository.MatchRepository, broadcast *service.BroadcastService) *MatchApplication {
	return &MatchApplication{
		engine:    engine,
		repo:      repo,
		broadcast: broadcast,
	}
}

func (app *MatchApplication) FindMatch(player *domain.Player, teamType domain.TeamType) *domain.Match {
	waitings := app.repo.WaitingMatches(teamType)

	var match domain.Match

	team := domain.NewTeam(teamType, player)
	if len(waitings) > 0 {
		match = *waitings[0]
		match.Join(&team)
	} else {
		match = domain.NewMatch(&team)
	}
	app.repo.Save(&match)

	if match.IsReady() {
		app.StartMatch(&match)
	}

	return &match
}

func (app *MatchApplication) StartMatch(match *domain.Match) {
	if !match.Start() {
		return
	}

	app.engine.NewGameLoop(match.ID)

	app.repo.Save(match)

	command := rpc.NewCommand("match/start", nil)
	app.broadcast.BroadcastToMatch(match, command)
}
