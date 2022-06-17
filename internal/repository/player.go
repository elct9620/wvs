package repository

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
)

type PlayerRepository struct {
	store *store.Store
}

func NewPlayerRepository(store *store.Store) *PlayerRepository {
	return &PlayerRepository{
		store: store,
	}
}

func (repo *PlayerRepository) Find(id string) (*domain.Player, error) {
	res, err := repo.store.Table("players").Find(id)
	if err != nil {
		return nil, errors.New("player not exists")
	}

	player := res.(domain.Player)

	return &player, nil
}

func (repo *PlayerRepository) Insert(player domain.Player) error {
	err := repo.store.Table("players").Insert(player.ID, player)
	if err != nil {
		return errors.New("player is exists")
	}

	return nil
}

func (repo *PlayerRepository) Delete(id string) {
	repo.store.Table("players").Delete(id)
}
