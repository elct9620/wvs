package repository

import (
	"errors"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository/schema"
)

var ErrPlayerNotFound = errors.New("player not found")
var ErrPlayerIDExists = errors.New("player id is exists")

type Players interface {
	Find(id string) (*domain.Player, error)
	Create(id string) error
	Delete(id string) error
}

type SimplePlayerRepository struct {
	items map[string]schema.Player
}

func NewSimplePlayerRepository() *SimplePlayerRepository {
	return &SimplePlayerRepository{
		items: make(map[string]schema.Player),
	}
}

func (repo *SimplePlayerRepository) Find(id string) (*domain.Player, error) {
	record, ok := repo.items[id]
	if ok == false {
		return nil, ErrPlayerNotFound
	}

	player := domain.NewPlayer(record.ID)
	return &player, nil
}

func (repo *SimplePlayerRepository) Create(id string) error {
	_, ok := repo.items[id]
	if ok {
		return ErrPlayerIDExists
	}

	repo.items[id] = schema.Player{ID: id}
	return nil
}

func (repo *SimplePlayerRepository) Delete(id string) error {
	delete(repo.items, id)
	return nil
}
