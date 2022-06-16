package hub

import (
	"errors"
	"sync"
)

type Publisher interface {
	WriteJSON(interface{}) error
}

type channel struct {
	sync.Mutex
	publisher Publisher
	messages  chan interface{}
	running   bool
	exit      chan bool
}

func (hub *Hub) NewChannel(id string, publisher Publisher) error {
	if _, ok := hub.channels[id]; ok == true {
		return errors.New("channel is exists")
	}

	hub.channels[id] = &channel{
		publisher: publisher,
		messages:  make(chan interface{}, 100),
		running:   false,
		exit:      make(chan bool),
	}

	return nil
}

func (hub *Hub) StartChannel(id string) error {
	channel, ok := hub.channels[id]
	if ok != true {
		return errors.New("channel not exists")
	}

	if channel.running == true {
		return errors.New("channel is running")
	}

	channel.Lock()
	channel.running = true
	go func() {
		for {
			select {
			case msg := <-channel.messages:
				channel.publisher.WriteJSON(msg)
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
		return errors.New("channel not exists")
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
