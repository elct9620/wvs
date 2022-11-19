package usecase

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
)

type Player struct {
	playerRepo *repository.PlayerRepository
}

func NewPlayer(playerRepo *repository.PlayerRepository) *Player {
	return &Player{
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
	app.playerRepo.Delete(id)
}
