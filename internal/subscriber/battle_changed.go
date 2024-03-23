package subscriber

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/usecase"
)

const (
	TopicBattleChanged = "battle"
	BattleChanged      = "battle_changed"
)

var _ Subscriber = &BattleChangedSubscriber{}

type BattleChangedSubscriber struct {
	startBattle *usecase.StartBattleCommand
}

func NewBattleChangedSubscriber(
	startBattle *usecase.StartBattleCommand,
) *BattleChangedSubscriber {
	return &BattleChangedSubscriber{
		startBattle: startBattle,
	}
}

func (s *BattleChangedSubscriber) Name() string {
	return BattleChanged
}

func (s *BattleChangedSubscriber) Topic() string {
	return TopicBattleChanged
}

func (s *BattleChangedSubscriber) Handler(msg *message.Message) error {
	var change DatabaseChange[db.BattleEvent]
	if err := json.Unmarshal(msg.Payload, &change); err != nil {
		return err
	}

	if change.Deleted() {
		return nil
	}

	id := change.After.AggregateId
	_, err := s.startBattle.Execute(context.Background(), &usecase.StartBattleInput{
		BattleId: id,
	})

	if err != nil {
		msg.Ack()
		return err
	}

	return nil
}
