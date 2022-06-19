package application

import (
	"time"

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
	gameLoop  *service.GameLoopService
}

func NewMatchApplication(engine *engine.Engine, repo *repository.MatchRepository, broadcast *service.BroadcastService, gameLoop *service.GameLoopService) *MatchApplication {
	return &MatchApplication{
		engine:    engine,
		repo:      repo,
		broadcast: broadcast,
		gameLoop:  gameLoop,
	}
}

func (app *MatchApplication) FindMatch(player *domain.Player, teamType domain.TeamType) (*domain.Match, bool) {
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
	app.repo.Save(&match)

	if match.IsMatched() {
		go func() {
			time.Sleep(10 * time.Millisecond)
			app.StartMatch(&match)
		}()
	}

	return &match, isTeam1
}

func (app *MatchApplication) StartMatch(match *domain.Match) {
	if !match.Start() {
		return
	}

	app.engine.NewGameLoop(match.ID, app.gameLoop.CreateLoop(match))
	app.repo.Save(match)

	command := rpc.NewCommand("match/start", nil)
	app.broadcast.BroadcastToMatch(match, command)
}
