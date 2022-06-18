package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/command/parameter"
)

type MatchApplication struct {
	BaseApplication
	repo *repository.MatchRepository
}

func NewMatchApplication(hub *hub.Hub, repo *repository.MatchRepository) *MatchApplication {
	return &MatchApplication{
		BaseApplication: BaseApplication{hub: hub},
		repo:            repo,
	}
}

func (app *MatchApplication) InitMatch(player *domain.Player, teamType domain.TeamType) error {
	team := domain.NewTeam(teamType, player)
	match := domain.NewMatch(&team)
	app.repo.Save(match)

	app.hub.PublishTo(player.ID, rpc.NewCommand("match/init", parameter.MatchInitParameter{ID: match.ID, Team: teamType}))
	return nil
}
