package usecase

import (
	"github.com/elct9620/wvs/internal/repository"
)

type Player struct {
	players repository.Players
}

func NewPlayer(players repository.Players) *Player {
	return &Player{
		players: players,
	}
}

func (usecase *Player) Register(sessionID string) error {
	err := usecase.players.Create(sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Player) Unregister(id string) {
	usecase.players.Delete(id)
}
