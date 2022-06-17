package application

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/utils"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
	"github.com/google/uuid"
)

type GameApplication struct {
	BaseApplication
}

func NewGameApplication(hub *hub.Hub) *GameApplication {
	return &GameApplication{
		BaseApplication: BaseApplication{hub: hub},
	}
}

func (app *GameApplication) ProcessCommand(player *domain.Player, command data.Command) error {
	evtName, err := utils.EventType(command)
	if err != nil {
		app.RaiseError(player, err.Error())
		return err
	}

	switch evtName {
	case "new":
		app.hub.PublishTo(player.ID, data.NewCommand("game", event.NewGameEvent{Room: uuid.NewString(), BaseEvent: event.BaseEvent{Type: "new"}}))
		return nil
	default:
		app.RaiseError(player, "unknown event")
		return errors.New("unknown event")
	}
}
