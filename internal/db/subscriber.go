package db

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/hashicorp/go-memdb"
)

var _ message.Subscriber = &Subscriber{}

type Subscriber struct {
	mux        sync.RWMutex
	watcher    *Watcher
	subscriber map[string][]chan *message.Message
	startOnce  sync.Once
}

func NewSubscriber(watcher *Watcher) *Subscriber {
	return &Subscriber{
		watcher:    watcher,
		subscriber: make(map[string][]chan *message.Message),
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.subscriber[topic]; !ok {
		s.subscriber[topic] = make([]chan *message.Message, 0)
	}

	out := make(chan *message.Message)
	s.subscriber[topic] = append(s.subscriber[topic], out)

	s.startOnce.Do(func() {
		go s.consume(ctx)
	})

	return out, nil
}

func (s *Subscriber) Close() error {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, subs := range s.subscriber {
		for _, sub := range subs {
			close(sub)
		}
	}

	s.watcher.Close()
	return nil
}

func (s *Subscriber) consume(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case change := <-s.watcher.ch:
			s.dispatch(change)
		}
	}
}

func (s *Subscriber) dispatch(change *memdb.Change) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if change == nil {
		return
	}

	subscribers, ok := s.subscriber[change.Table]
	if !ok {
		return
	}

	for _, out := range subscribers {
		event, err := databaseChangeToMessage(change)
		if err != nil {
			continue
		}

		out <- event
	}
}

func databaseChangeToMessage(change *memdb.Change) (*message.Message, error) {
	payload, err := json.Marshal(change)
	if err != nil {
		return nil, err
	}

	return message.NewMessage(
		watermill.NewUUID(),
		payload,
	), nil
}
