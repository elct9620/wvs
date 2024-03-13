package testability

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	ProvideRoutes,
	New,
)

type Route interface {
	Method() string
	Path() string
	http.Handler
}

type Testability struct {
	chi.Router
}

func New(routes ...Route) *Testability {
	router := chi.NewRouter()

	for _, route := range routes {
		router.Method(route.Method(), route.Path(), route)
	}

	return &Testability{
		Router: router,
	}
}

func ProvideRoutes() []Route {
	return []Route{}
}
