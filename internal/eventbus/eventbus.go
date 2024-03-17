package eventbus

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	New,
)

func New() (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, nil)
	if err != nil {
		return nil, err
	}

	router.AddPlugin(plugin.SignalsHandler)

	return router, nil
}