package subscriber

import (
	"github.com/ThreeDotsLabs/watermill/message"
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

func ProvideDatabaseSubscribers() []DatabaseSubscriber {
	return []DatabaseSubscriber{
		NewMatchChangedSubscriber(),
	}
}
