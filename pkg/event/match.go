package event

import "github.com/elct9620/wvs/internal/domain"

type InitMatchEvent struct {
	BaseEvent
	Team domain.Team `json:"team"`
}

type JoinMatchEvent struct {
	BaseEvent
	MatchID string `json:"match_id"`
}

func NewJoinMatchEvent(matchID string) JoinMatchEvent {
	return JoinMatchEvent{
		BaseEvent: BaseEvent{Type: "join"},
		MatchID:   matchID,
	}
}
