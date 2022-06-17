package application

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
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
	if command.Payload == nil {
		app.RaiseError(player, "invalid event")
		return errors.New("invalid event")
	}

	evt := command.Payload.(event.BaseEvent)
	switch evt.Type {
	case "new":
		app.hub.PublishTo(player.ID, data.NewCommand("game", event.NewGameEvent{Room: uuid.NewString(), BaseEvent: event.BaseEvent{Type: "new"}}))
		return nil
	default:
		app.RaiseError(player, "unknown event")
		return errors.New("unknown event")
	}
}
