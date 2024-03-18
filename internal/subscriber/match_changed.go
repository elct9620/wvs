package subscriber

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	TopicMatchChanged = "db:match"
	MatchChanged      = "match_changed"
)

var _ Subscriber = &MatchChangedSubscriber{}

type MatchChangedSubscriber struct {
}

func NewMatchChangedSubscriber() *MatchChangedSubscriber {
	return &MatchChangedSubscriber{}
}

func (s *MatchChangedSubscriber) Name() string {
	return MatchChanged
}

func (s *MatchChangedSubscriber) Topic() string {
	return TopicMatchChanged
}

func (s *MatchChangedSubscriber) Handler(msg *message.Message) error {
	return nil
}
