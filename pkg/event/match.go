package event

import (
	"encoding/json"
	"time"
)

var _ Event = &JoinMatchEvent{}

type JoinMatchEvent struct {
	matchId   string
	createdAt time.Time
	PlayerId  string `json:"player_id"`
}

func NewJoinMatchEvent(matchId string, playerId string) *JoinMatchEvent {
	return &JoinMatchEvent{
		matchId:   matchId,
		PlayerId:  playerId,
		createdAt: time.Now(),
	}
}

func (e *JoinMatchEvent) Type() string {
	return "JoinMatchEvent"
}

func (e *JoinMatchEvent) AggregateId() string {
	return e.matchId
}

func (e *JoinMatchEvent) CreatedAt() time.Time {
	return e.createdAt
}

func (e *JoinMatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":         e.Type(),
		"aggregate_id": e.AggregateId(),
		"created_at":   e.CreatedAt(),
		"player_id":    e.PlayerId,
	})
}

func (e *JoinMatchEvent) UnmarshalJSON(data []byte) error {
	var v struct {
		AggregateId string    `json:"aggregate_id"`
		CreatedAt   time.Time `json:"created_at"`
		PlayerId    string    `json:"player_id"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e.matchId = v.AggregateId
	e.createdAt = v.CreatedAt
	e.PlayerId = v.PlayerId
	return nil
}
