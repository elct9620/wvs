package app

import (
	"net/http"

	"github.com/elct9620/wvs/app/web"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	web.DefaultSet,
	New,
)

type Application struct {
	chi.Router
}

func New(
	web *web.Web,
) *Application {
	mux := chi.NewRouter()
	mux.Mount("/", web)
	return &Application{mux}
}

func (app *Application) Serve() error {
	return http.ListenAndServe(":8080", app)
}
