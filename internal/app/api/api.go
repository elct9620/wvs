package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	New,
)

type Api struct {
	chi.Router
}

func New() *Api {
	return &Api{
		Router: chi.NewRouter(),
	}
}
