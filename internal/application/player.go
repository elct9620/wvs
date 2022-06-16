package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/gorilla/websocket"
)

type PlayerApplication struct {
	playerRepo *repository.PlayerRepository
}

func NewPlayerApplication(playerRepo *repository.PlayerRepository) *PlayerApplication {
	return &PlayerApplication{
		playerRepo: playerRepo,
	}
}

func (app *PlayerApplication) Register(conn *websocket.Conn) (domain.Player, error) {
	player := domain.NewPlayerFromConn(conn)
	return player, app.playerRepo.Insert(player)
}

func (app *PlayerApplication) Unregister(id string) {
	app.playerRepo.Delete(id)
}
