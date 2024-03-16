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
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/repository/inmemory"
	"github.com/elct9620/wvs/internal/usecase"
)

// Injectors from wire.go:

func Initialize() (*app.Application, error) {
	scene := web.NewScene()
	webWeb := web.New(scene)
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	matchRepository := inmemory.NewMatchRepository(database)
	streamRepository := ws.NewStreamRepository()
	createMatchCommand := usecase.NewCreateMatchCommand(matchRepository, streamRepository)
	v := api.ProvideRoutes(createMatchCommand)
	apiApi := api.New(v...)
	subscribeCommand := usecase.NewSubscribeCommand(streamRepository)
	webSocket := ws.New(subscribeCommand, streamRepository)
	viper, err := config.NewViperWithDefaults()
	if err != nil {
		return nil, err
	}
	appConfig := app.NewConfig(viper)
	application := app.New(webWeb, apiApi, webSocket, appConfig)
	return application, nil
}
