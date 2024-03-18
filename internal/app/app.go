package app

import (
	"context"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/app/api"
	"github.com/elct9620/wvs/internal/app/web"
	"github.com/elct9620/wvs/internal/app/ws"
	"github.com/elct9620/wvs/internal/subscriber"
	"github.com/elct9620/wvs/internal/testability"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
)

var DefaultSet = wire.NewSet(
	web.DefaultSet,
	api.DefaultSet,
	ws.DefaultSet,
	subscriber.DefaultSet,
	ProvideEventBus,
	ProvideEventSubscribers,
	ProvideHttpServer,
	NewConfig,
	New,
)

var TestSet = wire.NewSet(
	web.DefaultSet,
	api.DefaultSet,
	ws.DefaultSet,
	testability.DefaultSet,
	subscriber.DefaultSet,
	ProvideEventBus,
	ProvideEventSubscribers,
	ProvideHttpTestServer,
	NewConfig,
	New,
)

type Application struct {
	chi.Router
	config *Config
	event  *message.Router
}

func New(
	http *chi.Mux,
	config *Config,
	event *message.Router,
) *Application {
	return &Application{
		Router: http,
		config: config,
		event:  event,
	}
}

func (app *Application) Serve() error {
	group := errgroup.Group{}

	group.Go(func() error {
		return app.event.Run(context.Background())
	})

	group.Go(func() error {
		return http.ListenAndServe(app.config.Address, app)
	})

	return group.Wait()
}

func (app *Application) StartEventBus() error {
	return app.event.Run(context.Background())
}

func (app *Application) StopEventBus() error {
	return app.event.Close()
}
