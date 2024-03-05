package app

import (
	"net/http"

	"github.com/elct9620/wvs/internal/app/api"
	"github.com/elct9620/wvs/internal/app/web"
	"github.com/elct9620/wvs/internal/app/ws"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	web.DefaultSet,
	api.DefaultSet,
	ws.DefaultSet,
	NewConfig,
	New,
)

type Application struct {
	chi.Router
	config *Config
}

func New(
	web *web.Web,
	api *api.Api,
	ws *ws.WebSocket,
	config *Config,
) *Application {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)

	mux.Mount("/", web)
	mux.Mount("/api", api)
	mux.Mount("/ws", ws)

	return &Application{mux, config}
}

func (app *Application) Serve() error {
	return http.ListenAndServe(app.config.Address, app)
}
