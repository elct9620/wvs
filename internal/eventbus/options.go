package eventbus

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

func WithDefaultMiddleware() RouterOptionFn {
	return func(router *message.Router) {
		router.AddMiddleware(
			middleware.CorrelationID,
			middleware.Retry{
				MaxRetries: 3,
			}.Middleware,
			middleware.Recoverer,
		)
	}
}
