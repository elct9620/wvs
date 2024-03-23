package battle

type handler func(*Battle, Event)

var handlers = map[string]handler{
	EventCreated: onCreated,
}

type Battle struct {
	id            string
	pendingEvents []Event
}

func New(id string) *Battle {
	battle := &Battle{}

	event := NewBattleCreated(id)
	battle.pendingEvents = append(battle.pendingEvents, event)
	battle.apply(event)

	return battle
}

func (b *Battle) Id() string {
	return b.id
}

func (b *Battle) PendingEvents() []Event {
	return b.pendingEvents
}

func (b *Battle) apply(evt Event) {
	handler, ok := handlers[evt.Type()]
	if !ok {
		return
	}

	handler(b, evt)
}

func onCreated(b *Battle, evt Event) {
	b.id = evt.(*BattleCreated).BattleId
}
