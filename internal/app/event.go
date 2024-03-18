package app

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/subscriber"
)

type EventBusOption = func(*message.Router)

func ProvideEventBus(options ...EventBusOption) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, nil)
	if err != nil {
		return nil, err
	}

	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries: 3,
		}.Middleware,
		middleware.Recoverer,
	)

	return router, nil
}

func ProvideEventSubscribers(
	database *db.Database,
	databaseSubscribers []subscriber.DatabaseSubscriber,
) []EventBusOption {
	watcher := database.Watch()
	dbSubscriber := db.NewSubscriber(watcher)

	return []EventBusOption{
		func(router *message.Router) {
			for _, s := range databaseSubscribers {
				router.AddNoPublisherHandler(s.Topic(), s.Name(), dbSubscriber, s.Handler)
			}
		},
	}
}
