package eventbus

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	ProvideOptions,
	New,
)

type RouterOptionFn func(*message.Router)

func New(options ...RouterOptionFn) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, nil)
	if err != nil {
		return nil, err
	}

	router.AddPlugin(plugin.SignalsHandler)

	for _, option := range options {
		option(router)
	}

	return router, nil
}

func ProvideOptions() []RouterOptionFn {
	return []RouterOptionFn{
		WithDefaultMiddleware(),
	}
}
