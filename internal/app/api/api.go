package api

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

type Api struct {
	chi.Router
}

func New(routes ...Route) *Api {
	router := chi.NewRouter()

	for _, route := range routes {
		router.Method(route.Method(), route.Path(), route)
	}

	return &Api{
		Router: router,
	}
}

func ProvideRoutes() []Route {
	return []Route{
		NewGetMe(),
	}
}
