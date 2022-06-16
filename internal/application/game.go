package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/data"
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
	app.hub.PublishTo(player.ID, command)
	return nil
}
