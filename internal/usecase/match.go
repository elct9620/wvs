package usecase

import (
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/rpc"
)

type Match struct {
	engine    *engine.Engine
	repo      *repository.MatchRepository
	broadcast *service.BroadcastService
	gameLoop  *service.GameLoopService
}

func NewMatch(engine *engine.Engine, repo *repository.MatchRepository, broadcast *service.BroadcastService, gameLoop *service.GameLoopService) *Match {
	return &Match{
		engine:    engine,
		repo:      repo,
		broadcast: broadcast,
		gameLoop:  gameLoop,
	}
}

func (app *Match) FindMatch(playerID string, teamType domain.TeamType) (*domain.Match, bool) {
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
	app.repo.Save(&match)

	if match.IsMatched() {
		go func() {
			time.Sleep(10 * time.Millisecond)
			app.StartMatch(&match)
		}()
	}

	return &match, isTeam1
}

func (app *Match) StartMatch(match *domain.Match) {
	if !match.Start() {
		return
	}

	app.engine.NewGameLoop(match.ID, app.gameLoop.CreateLoop(match))
	app.repo.Save(match)

	command := rpc.NewCommand("match/start", nil)
	app.broadcast.BroadcastToMatch(match, command)
}

func (app *Match) JoinMatch(matchID string, playerID string) bool {
	player := domain.Player{ID: playerID}
	match := app.repo.Find(matchID)
	if match == nil {
		return false
	}

	match.MarkReady(player.ID)
	app.repo.Save(match)

	if match.IsReady() {
		app.engine.StartGameLoop(match.ID)
	}

	return true
}