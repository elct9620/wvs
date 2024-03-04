package app

import (
	"net/http"

	"github.com/elct9620/wvs/app/api"
	"github.com/elct9620/wvs/app/web"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	web.DefaultSet,
	api.DefaultSet,
	New,
)

type Application struct {
	chi.Router
}

func New(
	web *web.Web,
	api *api.Api,
) *Application {
	mux := chi.NewRouter()

	mux.Mount("/", web)
	mux.Mount("/api", api)

	return &Application{mux}
}

func (app *Application) Serve() error {
	return http.ListenAndServe(":8080", app)
}
