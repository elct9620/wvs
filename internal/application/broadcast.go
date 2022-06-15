package application

import (
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/data"
)

type BroadcastApplication struct {
	playerRepo *repository.PlayerRepository
}

func NewBroadcastApplication(playerRepo *repository.PlayerRepository) *BroadcastApplication {
	return &BroadcastApplication{
		playerRepo: playerRepo,
	}
}

func (app *BroadcastApplication) BroadcastTo(playerID string, command data.Command) error {
	player, err := app.playerRepo.Find(playerID)
	if err != nil {
		return err
	}

	if player.Conn == nil {
		return nil
	}

	return player.Conn.WriteJSON(command)
}
