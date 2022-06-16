package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/gorilla/websocket"
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

func (app *PlayerApplication) Register(conn *websocket.Conn) (domain.Player, error) {
	player := domain.NewPlayer()
	err := app.playerRepo.Insert(player)
	if err != nil {
		return player, err
	}
	err = app.hub.NewChannel(player.ID, conn)
	return player, err
}

func (app *PlayerApplication) Unregister(id string) {
	app.hub.RemoveChannel(id)
	app.playerRepo.Delete(id)
}
