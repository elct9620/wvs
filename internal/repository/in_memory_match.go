package repository

import (
	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/internal/usecase"
)

var _ usecase.MatchRepository = &InMemoryMatchRepository{}

type InMemoryMatchRepository struct {
}

func NewInMemoryMatchRepository() *InMemoryMatchRepository {
	return &InMemoryMatchRepository{}
}

func (r *InMemoryMatchRepository) Save(match *match.Match) error {
	return nil
}
