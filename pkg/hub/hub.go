package hub

import (
	"context"
	"errors"
)

type Hub struct {
	channels map[string]*channel
	ctx      context.Context
	stop     context.CancelFunc
}

func NewHub() *Hub {
	ctx, stop := context.WithCancel(context.Background())

	return &Hub{
		channels: make(map[string]*channel),
		ctx:      ctx,
		stop:     stop,
	}
}

func (hub *Hub) PublishTo(channelID string, payload []byte) error {
	channel, ok := hub.channels[channelID]
	if ok != true {
		return errors.New("channel not exists")
	}

	channel.Lock()
	channel.messages <- payload
	channel.Unlock()

	return nil
}

func (hub *Hub) Stop() {
	hub.stop()
}
