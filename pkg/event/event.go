package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Payload   any       `json:"payload"`
}

const (
	EventReady = "ReadyEvent"
)

type ReadyEvent struct{}

func NewReadyEvent() *Event {
	return &Event{
		Id:        uuid.NewString(),
		Type:      EventReady,
		CreatedAt: time.Now(),
		Payload:   ReadyEvent{},
	}
}
