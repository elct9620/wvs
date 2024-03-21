package inmemory

import (
	"context"

	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/entity/battle"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/google/uuid"
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

	record := &db.BattleEvent{
		Id:      uuid.NewString(),
		MatchId: entity.Id(),
	}

	if err := txn.Insert(db.TableBattle, record); err != nil {
		return err
	}

	txn.Commit()
	return nil
}
