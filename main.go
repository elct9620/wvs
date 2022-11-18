package main

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/pkg/command"
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"golang.org/x/net/context"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			NewHub,
			NewEngine,
			NewStore,
			container.NewContainer,
			NewController,
		),
		fx.Invoke(func(*echo.Echo) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, controller *controller.WebSocketController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))
	e.GET("/ws", controller.Server)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					e.Logger.Fatal(e.Start(":8080"))
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return e.Shutdown(ctx)
			},
		},
	)

	return e
}

func NewController(container *container.Container) *controller.WebSocketController {
	playerRepo := container.NewPlayerRepository()

	service := command.NewRPCService(container)
	player := application.NewPlayerApplication(container.Hub(), playerRepo)
	controller := controller.NewWebSocketController(&service.RPC, container.Hub(), player)

	return controller
}

func NewHub(lc fx.Lifecycle) *hub.Hub {
	hub := hub.NewHub()

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			hub.Stop()
			return nil
		},
	})

	return hub
}

func NewStore() *store.Store {
	return infrastructure.InitStore()
}

func NewEngine(lc fx.Lifecycle) *engine.Engine {
	engine := engine.NewEngine()

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			engine.Stop()
			return nil
		},
	})

	return engine
}
