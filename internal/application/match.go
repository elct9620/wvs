package application

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/utils"
	"github.com/elct9620/wvs/pkg/data"
)

type MatchApplication struct {
	BaseApplication
}

func NewMatchApplication(hub *hub.Hub) *MatchApplication {
	return &MatchApplication{
		BaseApplication: BaseApplication{hub: hub},
	}
}

func (app *MatchApplication) ProcessCommand(player *domain.Player, command data.Command) error {
	evtName, err := utils.EventType(command)
	if err != nil {
		app.RaiseError(player, err.Error())
		return err
	}

	switch evtName {
	default:
		app.RaiseError(player, "unknown event")
		return errors.New("unknown event")
	}
}
