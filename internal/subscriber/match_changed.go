package subscriber

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/usecase"
)

const (
	TopicMatchChanged = "match"
	MatchChanged      = "match_changed"
)

var _ Subscriber = &MatchChangedSubscriber{}

type MatchChangedSubscriber struct {
	notifyJoinMatch *usecase.NotifyJoinMatchCommand
}

func NewMatchChangedSubscriber(notifyJoinMatch *usecase.NotifyJoinMatchCommand) *MatchChangedSubscriber {
	return &MatchChangedSubscriber{
		notifyJoinMatch: notifyJoinMatch,
	}
}

func (s *MatchChangedSubscriber) Name() string {
	return MatchChanged
}

func (s *MatchChangedSubscriber) Topic() string {
	return TopicMatchChanged
}

func (s *MatchChangedSubscriber) Handler(msg *message.Message) error {
	var change DatabaseChange[db.Match]
	if err := json.Unmarshal(msg.Payload, &change); err != nil {
		return err
	}

	var matchId string
	if change.After != nil {
		matchId = change.After.Id
	} else {
		matchId = change.Before.Id
	}

	_, err := s.notifyJoinMatch.Execute(msg.Context(), &usecase.NotifyJoinMatchInput{
		MatchId: matchId,
	})
	if err != nil {
		msg.Ack()
		return err
	}

	return nil
}
