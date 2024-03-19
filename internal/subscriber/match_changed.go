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

	isNewPlayerJoined := change.Created() || change.Updated()
	if !isNewPlayerJoined {
		return nil
	}

	for _, player := range change.After.Players {
		_, err := s.notifyJoinMatch.Execute(msg.Context(), &usecase.NotifyJoinMatchCommandInput{
			MatchId:  change.After.Id,
			PlayerId: player.Id,
		})

		if err != nil {
			msg.Ack()
			return err
		}
	}

	return nil
}
