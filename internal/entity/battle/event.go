package battle

import "github.com/google/uuid"

const (
	EventCreated = "Created"
)

type Event interface {
	Id() string
	AggregateId() string
	Type() string
	CreatedAt() int64
}

var _ Event = &BattleCreated{}

type BattleCreated struct {
	EventId   string
	BattleId  string
	createdAt int64
}

func NewBattleCreated(id string, createdAt int64) *BattleCreated {
	return &BattleCreated{
		EventId:   uuid.NewString(),
		BattleId:  id,
		createdAt: createdAt,
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

func (evt *BattleCreated) CreatedAt() int64 {
	return evt.createdAt
}
