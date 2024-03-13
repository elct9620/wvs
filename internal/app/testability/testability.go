package testability

import (
	"net/http"

	"github.com/elct9620/wvs/internal/usecase"
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

func ProvideRoutes(
	directCreateMatch usecase.Command[*usecase.DirectCreateMatchInput, *usecase.DirectCreateMatchOutput],
) []Route {
	return []Route{
		NewPostMatch(directCreateMatch),
	}
}
