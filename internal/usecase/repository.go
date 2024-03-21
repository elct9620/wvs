package usecase

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/battle"
	"github.com/elct9620/wvs/internal/entity/match"
)

type MatchRepository interface {
	Find(context.Context, string) (*match.Match, error)
	FindByPlayerID(context.Context, string) (*match.Match, error)
	Waiting(context.Context) ([]*match.Match, error)
	Save(context.Context, *match.Match) error
}

type BattleRepository interface {
	Save(context.Context, *battle.Battle) error
}
