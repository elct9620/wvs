//go:build embed
// +build embed

package server

import (
	"net/http"

	"go.uber.org/zap"
)

const isLiveReload = false

func WithAssets(logger *zap.Logger) HTTPOptionFn {
	return func(mux *http.ServeMux) {
	}
}
