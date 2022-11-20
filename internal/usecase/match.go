package usecase

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
)

type Match struct {
	engine   *engine.Engine
	repo     *repository.MatchRepository
	gameLoop *service.GameLoopService
}

func NewMatch(engine *engine.Engine, repo *repository.MatchRepository, gameLoop *service.GameLoopService) *Match {
	return &Match{
		engine:   engine,
		repo:     repo,
		gameLoop: gameLoop,
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
	app.repo.Save(&match)

	if match.IsMatched() {
		app.StartMatch(&match)
	}

	return &match, isTeam1, match.IsMatched()
}

func (app *Match) StartMatch(match *domain.Match) {
	if !match.Start() {
		return
	}

	app.engine.NewGameLoop(match.ID, app.gameLoop.CreateLoop(match))
	app.repo.Save(match)
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
