package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewScene,
	New,
)

type Web struct {
	chi.Router
}

func New(scene *Scene) *Web {
	mux := chi.NewRouter()
	mux.Method("GET", "/", scene)

	return &Web{mux}
}
