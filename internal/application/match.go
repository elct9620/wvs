package application

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/utils"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
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

func (app *MatchApplication) ProcessCommand(player *domain.Player, command data.Command) error {
	evtName, err := utils.EventType(command)
	if err != nil {
		app.RaiseError(player, err.Error())
		return err
	}

	switch evtName {
	case "init":
		return app.InitMatch(player, command.Payload.(event.InitMatchEvent))
	default:
		app.RaiseError(player, "unknown event")
		return errors.New("unknown event")
	}
}

func (app *MatchApplication) InitMatch(player *domain.Player, evt event.InitMatchEvent) error {
	team := domain.NewTeam(evt.Team, player)
	match := domain.NewMatch(&team)

	app.hub.PublishTo(player.ID, data.NewCommand("match", event.NewJoinMatchEvent(match.ID)))
	return nil
}
