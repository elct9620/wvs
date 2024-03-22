package event

import (
	"time"

	"github.com/google/uuid"
)

const (
	BattleStarted = "BattleStarted"
)

type BattleStartedEvent struct {
	Id string `json:"id"`
}

func NewBattleStartedEvent(id string) *Event {
	return &Event{
		Id:        uuid.NewString(),
		Type:      BattleStarted,
		CreatedAt: time.Now(),
		Payload: &BattleStartedEvent{
			Id: id,
		},
	}
}
