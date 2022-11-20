package hub

import (
	"errors"
	"sync"
)

var ErrChannelNotExists = errors.New("channel not exists")

type EventHandler func([]byte) error

type SimpleSubscriber struct {
	LastData string
}

func (p *SimpleSubscriber) OnEvent(v []byte) error {
	p.LastData = string(v)
	return nil
}

type channel struct {
	sync.Mutex
	handlers []EventHandler
	messages chan []byte
	running  bool
	exit     chan bool
}

func (hub *Hub) NewChannel(id string) error {
	if _, ok := hub.channels[id]; ok == true {
		return ErrChannelNotExists
	}

	hub.channels[id] = &channel{
		handlers: []EventHandler{},
		messages: make(chan []byte, 100),
		running:  false,
		exit:     make(chan bool),
	}

	return nil
}

func (hub *Hub) StartChannel(id string) error {
	channel, ok := hub.channels[id]
	if ok != true {
		return ErrChannelNotExists
	}

	if channel.running == true {
		return nil
	}

	channel.Lock()
	channel.running = true
	go func() {
		for {
			select {
			case event := <-channel.messages:
				for _, handler := range channel.handlers {
					handler(event)
				}
			case <-hub.ctx.Done():
				return
			case <-channel.exit:
				return
			}
		}
	}()
	channel.Unlock()
	return nil
}

func (hub *Hub) StopChannel(id string) error {
	channel, ok := hub.channels[id]
	if ok != true {
		return ErrChannelNotExists
	}

	if channel.running == false {
		return nil
	}

	channel.Lock()
	channel.running = false
	channel.exit <- true
	channel.Unlock()
	return nil
}

func (hub *Hub) RemoveChannel(id string) {
	hub.StopChannel(id)
	delete(hub.channels, id)
}

func (hub *Hub) AddHandler(id string, handler EventHandler) error {
	channel, ok := hub.channels[id]
	if ok != true {
		return ErrChannelNotExists
	}

	channel.handlers = append(channel.handlers, handler)
	return nil
}
