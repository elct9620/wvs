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

func (repo *InMemoryRooms) ListWaitings() ([]*entity.Room, error) {
	txn := repo.db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get(RoomTableName, "state", entity.RoomWaiting)
	if err != nil {
		return nil, err
	}
	rooms := []*entity.Room{}

	for row := it.Next(); row != nil; row = it.Next() {
		room := row.(*roomSchema)
		rooms = append(rooms, entity.NewRoom(
			room.ID,
			entity.WithRoomState(room.State),
		))
	}

	return rooms, nil
}

func (repo *InMemoryRooms) Save(room *entity.Room) error {
	txn := repo.db.Txn(true)
	defer txn.Commit()

	return txn.Insert(RoomTableName, &roomSchema{
		ID:    room.ID,
		State: room.State,
	})
}
