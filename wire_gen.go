// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wvs

import (
	"github.com/elct9620/wvs/internal/app"
	"github.com/elct9620/wvs/internal/app/api"
	"github.com/elct9620/wvs/internal/app/web"
	"github.com/elct9620/wvs/internal/app/ws"
	"github.com/elct9620/wvs/internal/config"
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/repository/inmemory"
	"github.com/elct9620/wvs/internal/testability"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/hashicorp/go-memdb"
)

// Injectors from wire.go:

func InitializeTest() (*app.Application, error) {
	scene := web.NewScene()
	webWeb := web.New(scene)
	dbSchema := db.ProvideDatabaseSchema()
	memDB, err := memdb.NewMemDB(dbSchema)
	if err != nil {
		return nil, err
	}
	matchRepository := inmemory.NewMatchRepository(memDB)
	streamRepository := ws.NewStreamRepository()
	createMatchCommand := usecase.NewCreateMatchCommand(matchRepository, streamRepository)
	v := api.ProvideRoutes(createMatchCommand)
	apiApi := api.New(v...)
	playerEventRepository := inmemory.NewPlayerEventRepository()
	subscribeCommand := usecase.NewSubscribeCommand(playerEventRepository, streamRepository)
	webSocket := ws.New(subscribeCommand, streamRepository)
	v2 := testability.ProvideRoutes(matchRepository)
	testabilityTestability := testability.New(v2...)
	viper, err := config.NewViperWithDefaults()
	if err != nil {
		return nil, err
	}
	appConfig := app.NewConfig(viper)
	application := app.NewTest(webWeb, apiApi, webSocket, testabilityTestability, appConfig)
	return application, nil
}
