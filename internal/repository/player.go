package repository

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
)

type PlayerRepository struct {
	store *store.Table
}

func NewPlayerRepository(store *store.Table) *PlayerRepository {
	return &PlayerRepository{
		store: store,
	}
}

func (repo *PlayerRepository) Find(id string) (*domain.Player, error) {
	res, err := repo.store.Find(id)
	if err != nil {
		return nil, errors.New("player not exists")
	}

	player := res.(domain.Player)

	return &player, nil
}

func (repo *PlayerRepository) Insert(player domain.Player) error {
	err := repo.store.Insert(player.ID, player)
	if err != nil {
		return errors.New("player is exists")
	}

	return nil
}

func (repo *PlayerRepository) Delete(id string) {
	repo.store.Delete(id)
}
