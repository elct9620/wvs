package db

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/hashicorp/go-memdb"
)

var _ message.Subscriber = &Subscriber{}

type Subscriber struct {
	mux        sync.RWMutex
	watcher    *Watcher
	subscriber map[string][]chan *message.Message
	runOnce    sync.Once
	closing    chan struct{}
	isClosed   bool
}

func NewSubscriber(watcher *Watcher) *Subscriber {
	return &Subscriber{
		watcher:    watcher,
		closing:    make(chan struct{}),
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

	s.runOnce.Do(func() {
		go s.consume(ctx)
	})

	go func() {
		<-s.closing
		close(out)
	}()

	return out, nil
}

func (s *Subscriber) Close() error {
	if s.isClosed {
		return nil
	}
	s.isClosed = true
	close(s.closing)

	return nil
}

func (s *Subscriber) consume(ctx context.Context) {
	for {
		select {
		case <-s.closing:
			return
		case <-ctx.Done():
			return
		case change := <-s.watcher.ch:
			s.dispatch(ctx, change)
		}
	}
}

func (s *Subscriber) dispatch(ctx context.Context, change *memdb.Change) {
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

		s.send(ctx, out, event)
	}
}

func (s *Subscriber) send(ctx context.Context, out chan *message.Message, msg *message.Message) {
	msgCtx, cancel := context.WithCancel(ctx)
	msg.SetContext(msgCtx)
	defer cancel()

ResendLoop:
	for {
		if s.isClosed {
			return
		}

		select {
		case out <- msg:
		case <-s.closing:
			return
		case <-ctx.Done():
			return
		}

		select {

		case <-msg.Acked():
			return
		case <-msg.Nacked():
			msg := msg.Copy()
			msg.SetContext(msgCtx)
			time.Sleep(100 * time.Millisecond)

			continue ResendLoop
		case <-s.closing:
			return
		case <-ctx.Done():
			return
		}
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
