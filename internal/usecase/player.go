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

func (app *Player) Register(sessionID string) error {
	player := domain.NewPlayer(sessionID)
	err := app.playerRepo.Insert(player)
	if err != nil {
		return err
	}

	return err
}

func (app *Player) Unregister(id string) {
	app.hub.RemoveChannel(id)
	app.playerRepo.Delete(id)
}
