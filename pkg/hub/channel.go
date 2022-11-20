package hub

import (
	"errors"
	"sync"
)

type Subscriber interface {
	OnEvent([]byte) error
}

type SimpleSubscriber struct {
	LastData string
}

func (p *SimpleSubscriber) OnEvent(v []byte) error {
	p.LastData = string(v)
	return nil
}

type channel struct {
	sync.Mutex
	subscriber Subscriber
	messages   chan []byte
	running    bool
	exit       chan bool
}

func (hub *Hub) NewChannel(id string, subscriber Subscriber) error {
	if _, ok := hub.channels[id]; ok == true {
		return errors.New("channel is exists")
	}

	hub.channels[id] = &channel{
		subscriber: subscriber,
		messages:   make(chan []byte, 100),
		running:    false,
		exit:       make(chan bool),
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
			case event := <-channel.messages:
				if channel.subscriber == nil {
					return
				}
				channel.subscriber.OnEvent(event)
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
		return errors.New("channel not exists")
	}

	if channel.running == false {
		return errors.New("channel not running")
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
