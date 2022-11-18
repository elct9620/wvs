package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/command/parameter"
	"github.com/elct9620/wvs/pkg/rpc"
)

type PlayerApplication struct {
	hub        *hub.Hub
	playerRepo *repository.PlayerRepository
}

func NewPlayerApplication(hub *hub.Hub, playerRepo *repository.PlayerRepository) *PlayerApplication {
	return &PlayerApplication{
		hub:        hub,
		playerRepo: playerRepo,
	}
}

func (app *PlayerApplication) Register(conn hub.Publisher) (domain.Player, error) {
	player := domain.NewPlayer()
	err := app.playerRepo.Insert(player)
	if err != nil {
		return player, err
	}

	err = app.hub.NewChannel(player.ID, conn)
	if err != nil {
		return player, err
	}

	err = app.hub.StartChannel(player.ID)
	if err != nil {
		return player, err
	}

	err = app.hub.PublishTo(player.ID, rpc.NewCommand("connected", parameter.ConnectedParameter{ID: player.ID}))
	return player, err
}

func (app *PlayerApplication) Unregister(id string) {
	app.hub.RemoveChannel(id)
	app.playerRepo.Delete(id)
}
