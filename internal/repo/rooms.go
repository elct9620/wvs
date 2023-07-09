package repository

import (
	"github.com/elct9620/wvs/internal/entity"
	"github.com/hashicorp/go-memdb"
)

const RoomTableName = "rooms"

type roomSchema struct {
	ID    string
	State int
}

type InMemoryRooms struct {
	db *memdb.MemDB
}

func NewInMemoryRoom(db *memdb.MemDB) *InMemoryRooms {
	return &InMemoryRooms{
		db: db,
	}
}

func (repo *InMemoryRooms) FindRoomBySessionID(id string) (*entity.Room, error) {
	txn := repo.db.Txn(false)
	defer txn.Abort()

	row, err := txn.First(PlayerTableName, "id", id)
	isPlayerNotExist := err != nil || row == nil
	if isPlayerNotExist {
		return nil, err
	}

	player := row.(*playerSchema)
	row, err = txn.First(RoomTableName, "id", player.RoomID)
	isRoomNotExist := err != nil || row == nil
	if isRoomNotExist {
		return nil, err
	}

	return buildRoomFromSchema(txn, row.(*roomSchema))
}

func (repo *InMemoryRooms) ListAvailable(team int) ([]*entity.Room, error) {
	txn := repo.db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get(RoomTableName, "state", entity.RoomWaiting)
	if err != nil {
		return nil, err
	}
	rooms := []*entity.Room{}

	for row := it.Next(); row != nil; row = it.Next() {
		room := row.(*roomSchema)
		entity, err := buildRoomFromSchema(txn, room)
		if err != nil {
			return nil, err
		}

		if entity.HasOpponent(team) {
			rooms = append(rooms, entity)
		}
	}

	return rooms, nil
}

func (repo *InMemoryRooms) Save(room *entity.Room) error {
	txn := repo.db.Txn(true)
	defer txn.Commit()

	err := txn.Insert(RoomTableName, buildRoomSchema(room))
	if err != nil {
		txn.Abort()
		return err
	}

	for _, player := range room.Players {
		err := txn.Insert(PlayerTableName, buildPlayerSchema(room.ID, player))

		if err != nil {
			txn.Abort()
			return err
		}
	}

	return nil
}

func buildRoomSchema(room *entity.Room) *roomSchema {
	return &roomSchema{
		ID:    room.ID,
		State: room.State,
	}
}

func buildRoomFromSchema(txn *memdb.Txn, roomData *roomSchema) (*entity.Room, error) {
	room := entity.NewRoom(roomData.ID, entity.WithRoomState(roomData.State))

	playerIt, err := txn.Get(PlayerTableName, "roomID", room.ID)
	if err != nil {
		return nil, err
	}

	for row := playerIt.Next(); row != nil; row = playerIt.Next() {
		playerData := row.(*playerSchema)
		err := room.AddPlayer(playerData.ID, playerData.Team)
		if err != nil {
			return nil, err
		}
	}

	return room, nil
}
