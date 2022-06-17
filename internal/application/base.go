package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
)

type BaseApplication struct {
	hub *hub.Hub
}

func NewBaseApplication(hub *hub.Hub) *BaseApplication {
	return &BaseApplication{hub: hub}
}

func (app *BaseApplication) RaiseError(player *domain.Player, reason string) {
	app.hub.PublishTo(player.ID, data.NewCommand("error", event.ErrorEvent{Message: reason}))
}
