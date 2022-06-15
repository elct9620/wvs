package repository

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
)

// In-memory store
var (
	players map[string]domain.Player = make(map[string]domain.Player)
)

type PlayerRepository struct {
	players *map[string]domain.Player
}

func NewPlayerRepository() PlayerRepository {
	return PlayerRepository{
		players: &players,
	}
}

func (repo *PlayerRepository) Insert(player domain.Player) error {
	if _, ok := (*repo.players)[player.ID]; ok == true {
		return errors.New("player is exists")
	}

	(*repo.players)[player.ID] = player

	return nil
}

func (repo *PlayerRepository) Delete(id string) error {
	if _, ok := (*repo.players)[id]; ok == true {
		delete(*repo.players, id)
	}

	return nil
}
