package battle

import (
	"sync"
	"time"
)

type handler func(*Battle, Event)

var handlers = map[string]handler{
	EventCreated: onCreated,
}

type Battle struct {
	mux           sync.RWMutex
	id            string
	pendingEvents []Event
}

func New(id string) *Battle {
	battle := &Battle{}

	event := NewBattleCreated(id, time.Now().Unix())
	battle.apply(event)

	return battle
}

func (b *Battle) Id() string {
	return b.id
}

func (b *Battle) PendingEvents() []Event {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.pendingEvents
}

func (b *Battle) ClearEvents() {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.pendingEvents = make([]Event, 0)
}

func (b *Battle) apply(evt Event) {
	b.mux.Lock()
	defer b.mux.Unlock()

	handler, ok := handlers[evt.Type()]
	if !ok {
		return
	}

	handler(b, evt)

	b.pendingEvents = append(b.pendingEvents, evt)
}

func onCreated(b *Battle, evt Event) {
	b.id = evt.(*BattleCreated).BattleId
}
