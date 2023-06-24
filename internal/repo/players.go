package repository

import (
	"github.com/elct9620/wvs/internal/entity"
	"github.com/hashicorp/go-memdb"
)

const PlayerTableName = "players"

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
		return row.(*entity.Player)
	}

	player := entity.NewPlayer(id)
	err = txn.Insert(PlayerTableName, player)
	if err != nil {
		return nil
	}

	return player
}

func (repo *InMemoryPlayers) Save(player *entity.Player) error {
	txn := repo.db.Txn(true)
	defer txn.Commit()

	return txn.Insert(PlayerTableName, player)
}
