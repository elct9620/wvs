package repository

import (
	"github.com/elct9620/wvs/internal/entity"
)

const PlayerTableName = "players"

type playerSchema struct {
	ID     string
	Team   int
	RoomID string
}

func buildPlayerSchema(roomID string, player *entity.Player) *playerSchema {
	return &playerSchema{
		ID:     player.ID,
		Team:   player.Team,
		RoomID: roomID,
	}
}
