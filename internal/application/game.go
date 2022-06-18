package application

import (
	"github.com/elct9620/wvs/internal/infrastructure/hub"
)

type GameApplication struct {
	BaseApplication
}

func NewGameApplication(hub *hub.Hub) *GameApplication {
	return &GameApplication{
		BaseApplication: BaseApplication{hub: hub},
	}
}
