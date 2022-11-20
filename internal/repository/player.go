package repository

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
)

var ErrPlayerNotFound = errors.New("player not found")
var ErrPlayerIDExists = errors.New("player id is exists")

type Players interface {
	Find(id string) (*domain.Player, error)
	Create(id string) error
	Delete(id string) error
}

type PlayerRecord struct {
	ID string
}

type SimplePlayerRepository struct {
	players map[string]PlayerRecord
}

func NewSimplePlayerRepository() *SimplePlayerRepository {
	return &SimplePlayerRepository{
		players: make(map[string]PlayerRecord),
	}
}

func (repo *SimplePlayerRepository) Find(id string) (*domain.Player, error) {
	record, ok := repo.players[id]
	if ok == false {
		return nil, ErrPlayerNotFound
	}

	player := domain.NewPlayer(record.ID)
	return &player, nil
}

func (repo *SimplePlayerRepository) Create(id string) error {
	_, ok := repo.players[id]
	if ok {
		return ErrPlayerIDExists
	}

	repo.players[id] = PlayerRecord{id}
	return nil
}

func (repo *SimplePlayerRepository) Delete(id string) error {
	delete(repo.players, id)
	return nil
}
