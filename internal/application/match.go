package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/repository"
)

type MatchApplication struct {
	hub  *hub.Hub
	repo *repository.MatchRepository
}

func NewMatchApplication(hub *hub.Hub, repo *repository.MatchRepository) *MatchApplication {
	return &MatchApplication{
		hub:  hub,
		repo: repo,
	}
}

func (app *MatchApplication) StartMatch(player *domain.Player, teamType domain.TeamType) *domain.Match {
	team := domain.NewTeam(teamType, player)
	match := domain.NewMatch(&team)
	app.repo.Save(match)

	return &match
}
