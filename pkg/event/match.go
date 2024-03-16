package event

import (
	"time"

	"github.com/google/uuid"
)

const (
	EventJoinMatch = "JoinMatchEvent"
)

type JoinMatchEvent struct {
	AggregateId string `json:"aggregate_id"`
	PlayerId    string `json:"player_id"`
}

func NewJoinMatchEvent(matchId string, playerId string) *Event {
	return &Event{
		Id:        uuid.NewString(),
		Type:      EventJoinMatch,
		CreatedAt: time.Now(),
		Payload: &JoinMatchEvent{
			AggregateId: matchId,
			PlayerId:    playerId,
		},
	}
}
