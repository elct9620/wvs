package event

import "github.com/elct9620/wvs/internal/domain"

type StartMatchEvent struct {
	BaseEvent
	Team domain.Team `json:"team"`
}
