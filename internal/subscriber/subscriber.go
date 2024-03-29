package subscriber

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	ProvideDatabaseSubscribers,
)

type Subscriber interface {
	Name() string
	Topic() string
	Handler(*message.Message) error
}

type DatabaseSubscriber Subscriber

func ProvideDatabaseSubscribers(
	notifyJoinMatch *usecase.CreateBattleCommand,
	startBattle *usecase.StartBattleCommand,
) []DatabaseSubscriber {
	return []DatabaseSubscriber{
		NewMatchChangedSubscriber(notifyJoinMatch),
		NewBattleChangedSubscriber(startBattle),
	}
}
