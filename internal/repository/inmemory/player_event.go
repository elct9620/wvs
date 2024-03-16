package inmemory

import (
	"context"
	"sync"

	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/event"
)

type playerEventWatcher struct {
	isWatching bool
	channel    chan event.Event
}

var _ usecase.PlayerEventRepository = &PlayerEventRepository{}

type PlayerEventRepository struct {
	mux      sync.RWMutex
	watchers map[string]*playerEventWatcher
}

func NewPlayerEventRepository() *PlayerEventRepository {
	return &PlayerEventRepository{
		watchers: make(map[string]*playerEventWatcher),
	}
}

func (r *PlayerEventRepository) Publish(ctx context.Context, sessionId string, event event.Event) error {
	r.mux.RLock()
	defer r.mux.RUnlock()

	watcher, isFound := r.watchers[sessionId]
	if !isFound {
		return nil
	}

	watcher.channel <- event
	return nil
}

func (r *PlayerEventRepository) Watch(ctx context.Context, sessionId string) (chan event.Event, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	watcher, isFound := r.watchers[sessionId]
	isUsed := isFound && watcher.isWatching
	if isUsed {
		return nil, usecase.ErrAlreadyWatching
	}

	if !isFound {
		watcher = &playerEventWatcher{
			channel: make(chan event.Event),
		}
	}
	watcher.isWatching = true
	r.watchers[sessionId] = watcher

	return watcher.channel, nil
}
