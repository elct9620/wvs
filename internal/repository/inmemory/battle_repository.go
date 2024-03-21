package inmemory

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/battle"
	"github.com/elct9620/wvs/internal/usecase"
)

var _ usecase.BattleRepository = &BattleRepository{}

type BattleRepository struct {
}

func NewBattleRepository() *BattleRepository {
	return &BattleRepository{}
}

func (r *BattleRepository) Save(ctx context.Context, entity *battle.Battle) error {
	return nil
}
