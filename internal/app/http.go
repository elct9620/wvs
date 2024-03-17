package app

import (
	"context"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/app/api"
	"github.com/elct9620/wvs/internal/app/web"
	"github.com/elct9620/wvs/internal/app/ws"
	"github.com/elct9620/wvs/internal/testability"
	"github.com/elct9620/wvs/pkg/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
)

var DefaultSet = wire.NewSet(
	web.DefaultSet,
	api.DefaultSet,
	ws.DefaultSet,
	NewConfig,
	New,
)

var TestSet = wire.NewSet(
	web.DefaultSet,
	api.DefaultSet,
	ws.DefaultSet,
	testability.DefaultSet,
	NewConfig,
	NewTest,
)

type Application struct {
	chi.Router
	eventBus *message.Router
	config   *Config
}

func New(
	web *web.Web,
	api *api.Api,
	ws *ws.WebSocket,
	config *Config,
	eventBus *message.Router,
) *Application {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)
	mux.Use(session.Middleware(config.SessionKey))

	mux.Mount("/", web)
	mux.Mount("/api", api)
	mux.Mount("/ws", ws)

	return &Application{mux, eventBus, config}
}

func NewTest(
	web *web.Web,
	api *api.Api,
	ws *ws.WebSocket,
	testability *testability.Testability,
	config *Config,
	eventBus *message.Router,
) *Application {
	mux := chi.NewRouter()

	mux.Use(session.Middleware(config.SessionKey))

	mux.Mount("/", web)
	mux.Mount("/api", api)
	mux.Mount("/ws", ws)
	mux.Mount("/testability", testability)

	return &Application{mux, eventBus, config}
}

func (app *Application) Serve() error {
	group := errgroup.Group{}

	group.Go(func() error {
		return app.eventBus.Run(context.Background())
	})

	group.Go(func() error {
		return http.ListenAndServe(app.config.Address, app)
	})

	return group.Wait()
}
