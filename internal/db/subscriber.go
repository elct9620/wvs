package db

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

var _ message.Subscriber = &Subscriber{}

type Subscriber struct {
	watcher *Watcher
}

func NewSubscriber(watcher *Watcher) *Subscriber {
	return &Subscriber{
		watcher: watcher,
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	ch := make(chan *message.Message)

	go func() {
		defer close(ch)

		for {
			change := s.watcher.Consume()
			if change == nil {
				continue
			}

			msg := message.NewMessage(watermill.NewUUID(), []byte(change.Table))
			ch <- msg
		}
	}()

	return ch, nil
}

func (s *Subscriber) Close() error {
	s.watcher.Close()

	return nil
}
