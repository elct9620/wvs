package usecase

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/command/parameter"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
)

type Player struct {
	hub        *hub.Hub
	playerRepo *repository.PlayerRepository
}

func NewPlayer(hub *hub.Hub, playerRepo *repository.PlayerRepository) *Player {
	return &Player{
		hub:        hub,
		playerRepo: playerRepo,
	}
}

func (app *Player) Register(conn hub.Subscriber) (string, error) {
	player := domain.NewPlayer()
	err := app.playerRepo.Insert(player)
	if err != nil {
		return player.ID, err
	}

	err = app.hub.NewChannel(player.ID, conn)
	if err != nil {
		return player.ID, err
	}

	err = app.hub.StartChannel(player.ID)
	if err != nil {
		return player.ID, err
	}

	err = app.hub.PublishTo(player.ID, rpc.NewCommand("connected", parameter.ConnectedParameter{ID: player.ID}))
	if err != nil {
		return player.ID, err
	}

	return player.ID, err
}

func (app *Player) Unregister(id string) {
	app.hub.RemoveChannel(id)
	app.playerRepo.Delete(id)
}
