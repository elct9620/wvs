package event

import (
	"encoding/json"
	"time"
)

type Event interface {
	json.Marshaler
	json.Unmarshaler
	Type() string
	AggregateId() string
	CreatedAt() time.Time
}

var _ Event = &ReadyEvent{}

type ReadyEvent struct {
	aggregateId string
	createdAt   time.Time
}

func NewReadyEvent(aggregateId string) *ReadyEvent {
	return &ReadyEvent{aggregateId: aggregateId, createdAt: time.Now()}
}

func (e *ReadyEvent) Type() string {
	return "ReadyEvent"
}

func (e *ReadyEvent) AggregateId() string {
	return e.aggregateId
}

func (e *ReadyEvent) CreatedAt() time.Time {
	return e.createdAt
}

func (e *ReadyEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":         e.Type(),
		"aggregate_id": e.AggregateId(),
		"created_at":   e.CreatedAt(),
	})
}

func (e *ReadyEvent) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	e.aggregateId = m["aggregate_id"].(string)
	e.createdAt = m["created_at"].(time.Time)
	return nil
}
