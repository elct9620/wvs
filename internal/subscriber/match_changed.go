package subscriber

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	TopicMatchChanged = "db:match"
)

func SubscribeMatchChanged(subscriber message.Subscriber) SubscriberFn {
	return func(r *message.Router) {
		r.AddNoPublisherHandler(
			TopicMatchChanged,
			"match_changed",
			subscriber,
			MatchChangedHandler,
		)
	}
}

func MatchChangedHandler(msg *message.Message) error {
	return nil
}
