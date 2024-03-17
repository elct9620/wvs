package app

import (
	"github.com/elct9620/wvs/internal/app/api"
	"github.com/elct9620/wvs/internal/app/web"
	"github.com/elct9620/wvs/internal/app/ws"
	"github.com/elct9620/wvs/internal/testability"
	"github.com/elct9620/wvs/pkg/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ProvideHttpServer(
	web *web.Web,
	api *api.Api,
	ws *ws.WebSocket,
	config *Config,

) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)
	mux.Use(session.Middleware(config.SessionKey))

	mux.Mount("/", web)
	mux.Mount("/api", api)
	mux.Mount("/ws", ws)

	return mux
}

func ProvideHttpTestServer(
	web *web.Web,
	api *api.Api,
	ws *ws.WebSocket,
	testability *testability.Testability,
	config *Config,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(session.Middleware(config.SessionKey))

	mux.Mount("/", web)
	mux.Mount("/api", api)
	mux.Mount("/ws", ws)
	mux.Mount("/testability", testability)

	return mux
}
