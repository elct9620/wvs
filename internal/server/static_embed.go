//go:build embed
// +build embed

package server

import (
	"net/http"

	"go.uber.org/zap"
)

func WithRoot(sessions SessionStore, logger *zap.Logger) HTTPOptionFn {
	return func(mux *http.ServeMux) {
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			sessions.Renew(req)

			w.WriteHeader(http.StatusOK)
		})
	}
}
