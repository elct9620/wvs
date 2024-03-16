package inmemory

import (
	"context"

	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/event"
)

var _ usecase.PlayerEventRepository = &PlayerEventRepository{}

type PlayerEventRepository struct {
}

func NewPlayerEventRepository() *PlayerEventRepository {
	return &PlayerEventRepository{}
}

func (r *PlayerEventRepository) Watch(ctx context.Context, sessionId string) (chan event.Event, error) {
	return make(chan event.Event), nil
}
