//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	controller "github.com/elct9620/wvs/internal/ctrl"
	"github.com/elct9620/wvs/internal/server"
	"github.com/elct9620/wvs/internal/usecases"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func initServer(logger *zap.Logger) (*http.ServeMux, error) {
	wire.Build(
		usecases.ProviderSet,
		controller.ProviderSet,
		server.ProvideInMemorySession,
		server.NewServices,
		server.NewRPC,
		server.NewRoutes,
		server.NewMux,
	)

	return nil, nil
}
