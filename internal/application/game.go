package application

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
)

type GameApplication struct {
	hub *hub.Hub
}

func NewGameApplication(hub *hub.Hub) *GameApplication {
	return &GameApplication{
		hub: hub,
	}
}

func (app *GameApplication) ProcessCommand(player *domain.Player, command data.Command) error {
	if command.Payload == nil {
		app.raiseError(player, "invalid event")
		return errors.New("invalid event")
	}

	evt := command.Payload.(event.BaseEvent)
	switch evt.Type {
	default:
		app.raiseError(player, "unknown event")
		return errors.New("unknown event")
	}
}

func (app *GameApplication) raiseError(player *domain.Player, reason string) {
	app.hub.PublishTo(player.ID, data.NewCommand("error", event.ErrorEvent{Message: reason}))
}
