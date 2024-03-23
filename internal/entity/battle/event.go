package battle

import "github.com/google/uuid"

const (
	EventCreated = "Created"
)

type Event interface {
	Id() string
	AggregateId() string
	Type() string
}

var _ Event = &BattleCreated{}

type BattleCreated struct {
	EventId  string
	BattleId string
}

func NewBattleCreated(id string) *BattleCreated {
	return &BattleCreated{
		EventId:  uuid.NewString(),
		BattleId: id,
	}
}

func (evt *BattleCreated) Id() string {
	return evt.EventId
}

func (evt *BattleCreated) AggregateId() string {
	return evt.BattleId
}

func (evt *BattleCreated) Type() string {
	return EventCreated
}
