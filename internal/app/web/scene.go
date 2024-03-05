package web

import "net/http"

var _ http.Handler = &Scene{}

type Scene struct {
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
