package usecase

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/hub"
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

	return player.ID, err
}

func (app *Player) Unregister(id string) {
	app.playerRepo.Delete(id)
}
