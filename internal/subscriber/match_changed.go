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
	createBattle *usecase.CreateBattleCommand
}

func NewMatchChangedSubscriber(createBattle *usecase.CreateBattleCommand) *MatchChangedSubscriber {
	return &MatchChangedSubscriber{
		createBattle: createBattle,
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

	if !change.Updated() {
		return nil
	}

	_, err := s.createBattle.Execute(msg.Context(), &usecase.CreateBattleInput{
		MatchId: change.After.Id,
	})
	if err != nil {
		msg.Ack()
		return err
	}

	return nil
}
