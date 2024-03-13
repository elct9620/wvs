package usecase

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/match"
)

type MatchRepository interface {
	FindByPlayerID(context.Context, string) (*match.Match, error)
	WaitingList(context.Context) ([]*match.Match, error)
	Save(context.Context, *match.Match) error
}
