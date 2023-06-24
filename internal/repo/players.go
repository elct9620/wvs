package repository

import (
	"github.com/elct9620/wvs/internal/entity"
)

const PlayerTableName = "players"

type playerSchema struct {
	ID     string
	RoomID string
}

func buildPlayerSchema(roomID string, player *entity.Player) *playerSchema {
	return &playerSchema{
		ID:     player.ID,
		RoomID: roomID,
	}
}

func buildPlayerFromSchema(player *playerSchema) *entity.Player {
	return entity.NewPlayer(
		player.ID,
	)
}
