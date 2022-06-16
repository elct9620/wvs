package domain

import (
	"github.com/google/uuid"
)

type Player struct {
	ID string
}

func NewPlayer() Player {
	return Player{
		ID: uuid.NewString(),
	}
}
