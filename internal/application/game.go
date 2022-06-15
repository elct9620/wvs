package application

import "github.com/elct9620/wvs/pkg/data"

type GameApplication struct {
}

func NewGameApplication() *GameApplication {
	return &GameApplication{}
}

func (app *GameApplication) StartGame(playerID string) (data.BroadcastTarget, data.Command, error) {
	return data.NewBroadcastTarget(false, playerID), data.NewCommand("event"), nil
}
