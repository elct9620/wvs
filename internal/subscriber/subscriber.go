package subscriber

import "github.com/ThreeDotsLabs/watermill/message"

type SubscriberFn func(r *message.Router)
