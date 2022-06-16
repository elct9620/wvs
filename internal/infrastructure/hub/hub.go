package hub

import "errors"

type Hub struct {
	channels map[string]*channel
}

func NewHub() *Hub {
	return &Hub{
		channels: make(map[string]*channel),
	}
}

func (hub *Hub) PublishTo(channelID string, data interface{}) error {
	channel, ok := hub.channels[channelID]
	if ok != true {
		return errors.New("channel not exists")
	}

	channel.Lock()
	channel.messages <- data
	channel.Unlock()

	return nil
}
