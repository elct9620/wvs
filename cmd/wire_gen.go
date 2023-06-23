// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/elct9620/wvs/internal/ctrl"
	"github.com/elct9620/wvs/internal/server"
	"go.uber.org/zap"
	"net/http"
)

// Injectors from wire.go:

func initServer(logger *zap.Logger) (*http.ServeMux, error) {
	system := controller.NewSystem()
	lobby := controller.NewLobby()
	v := server.NewServices(system, lobby)
	rpcServer, err := server.NewRPC(v...)
	if err != nil {
		return nil, err
	}
	inMemorySession := server.NewInMemorySession()
	v2 := server.NewRoutes(rpcServer, inMemorySession, logger)
	serveMux := server.NewMux(v2...)
	return serveMux, nil
}
