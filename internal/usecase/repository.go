package usecase

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/pkg/event"
)

type MatchRepository interface {
	FindByPlayerID(context.Context, string) (*match.Match, error)
	Waiting(context.Context) ([]*match.Match, error)
	Save(context.Context, *match.Match) error
}

type PlayerEventRepository interface {
	Watch(context.Context, string) (chan event.Event, error)
	Publish(context.Context, string, event.Event) error
}
