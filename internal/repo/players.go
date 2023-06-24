package repository

import (
	"github.com/elct9620/wvs/internal/entity"
	"github.com/hashicorp/go-memdb"
)

const PlayerTableName = "players"

type playerSchema struct {
	ID     string
	RoomID string
}

type InMemoryPlayers struct {
	db *memdb.MemDB
}

func NewInMemoryPlayer(db *memdb.MemDB) *InMemoryPlayers {
	return &InMemoryPlayers{
		db: db,
	}
}

func (repo *InMemoryPlayers) FindOrCreate(id string) *entity.Player {
	txn := repo.db.Txn(true)
	defer txn.Commit()

	row, err := txn.First(PlayerTableName, "id", id)
	if err != nil {
		return nil
	}

	if row != nil {
		player := row.(*playerSchema)
		return buildPlayerFromSchema(txn, player)
	}

	player := entity.NewPlayer(id)
	err = txn.Insert(PlayerTableName, &playerSchema{
		ID: player.ID,
	})
	if err != nil {
		return nil
	}

	return player
}

func (repo *InMemoryPlayers) Save(player *entity.Player) error {
	txn := repo.db.Txn(true)
	defer txn.Commit()

	var roomID string
	if player.Room != nil {
		roomID = player.Room.ID
	}

	return txn.Insert(PlayerTableName, &playerSchema{
		ID:     player.ID,
		RoomID: roomID,
	})
}

func buildPlayerFromSchema(txn *memdb.Txn, player *playerSchema) *entity.Player {
	options := make([]entity.PlayerOptionFn, 0)

	if len(player.RoomID) > 0 {
		room, err := txn.First(RoomTableName, "id", player.RoomID)
		if err != nil {
			return nil
		}

		options = append(options, entity.WithPlayerRoom(buildRoomFromSchema(txn, room.(*roomSchema))))
	}

	return entity.NewPlayer(
		player.ID,
		options...,
	)
}
