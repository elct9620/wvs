package eventbus

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/db"
)

func SubscribeDatabaseChanges(
	database *db.Database,
	handler *DatabaseChangeHandler,
) RouterOptionFn {
	subscriber := db.NewSubscriber(database.Watch())

	return func(r *message.Router) {
		r.AddNoPublisherHandler(
			"db_change_handler",
			"db_change",
			subscriber,
			handler.Handle,
		)
	}
}

type DatabaseChangeHandler struct {
}

func NewDatabaseChangeHandler() *DatabaseChangeHandler {
	return &DatabaseChangeHandler{}
}

func (h *DatabaseChangeHandler) Handle(msg *message.Message) error {
	return nil
}
