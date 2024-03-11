// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/elct9620/wvs/internal/app"
	"github.com/elct9620/wvs/internal/app/api"
	"github.com/elct9620/wvs/internal/app/web"
	"github.com/elct9620/wvs/internal/app/ws"
	"github.com/elct9620/wvs/internal/config"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/usecase"
)

// Injectors from wire.go:

func Initialize() (*app.Application, error) {
	scene := web.NewScene()
	webWeb := web.New(scene)
	inMemoryMatchRepository := repository.NewInMemoryMatchRepository()
	createMatchCommand := usecase.NewCreateMatchCommand(inMemoryMatchRepository)
	v := api.ProvideRoutes(createMatchCommand)
	apiApi := api.New(v...)
	webSocket := ws.New()
	viper, err := config.NewViperWithDefaults()
	if err != nil {
		return nil, err
	}
	appConfig := app.NewConfig(viper)
	application := app.New(webWeb, apiApi, webSocket, appConfig)
	return application, nil
}
