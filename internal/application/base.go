package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/pkg/command"
)

type BaseApplication struct {
	hub *hub.Hub
}

func NewBaseApplication(hub *hub.Hub) *BaseApplication {
	return &BaseApplication{hub: hub}
}

func (app *BaseApplication) RaiseError(player *domain.Player, reason string) {
	app.hub.PublishTo(player.ID, rpc.NewCommand("error", command.ErrorParameter{Reason: reason}))
}
