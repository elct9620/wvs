package repository

import (
	"github.com/elct9620/wvs/internal/entity"
	"github.com/hashicorp/go-memdb"
)

const RoomTableName = "rooms"

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

	it, err := txn.Get(RoomTableName, "id")
	if err != nil {
		return nil, err
	}
	rooms := []*entity.Room{}

	for row := it.Next(); row != nil; row = it.Next() {
		room := row.(entity.Room)
		rooms = append(rooms, &room)
	}

	return rooms, nil
}
