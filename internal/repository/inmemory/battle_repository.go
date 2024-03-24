package inmemory

import (
	"context"

	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/entity/battle"
	"github.com/elct9620/wvs/internal/usecase"
)

var _ usecase.BattleRepository = &BattleRepository{}

type BattleRepository struct {
	db *db.Database
}

func NewBattleRepository(db *db.Database) *BattleRepository {
	return &BattleRepository{
		db: db,
	}
}

func (r *BattleRepository) Save(ctx context.Context, entity *battle.Battle) error {
	txn := r.db.Txn(true)
	defer txn.Abort()

	events := entity.PendingEvents()

	for _, evt := range events {
		record := &db.BattleEvent{
			Id:          evt.Id(),
			AggregateId: evt.AggregateId(),
			Type:        evt.Type(),
		}

		if err := txn.Insert(db.TableBattle, record); err != nil {
			return err
		}
	}

	txn.Commit()
	entity.ClearEvents()

	return nil
}
