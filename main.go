package main

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/container"
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

func NewController(lc fx.Lifecycle) *controller.WebSocketController {
	container := container.NewContainer()
	playerRepo := container.NewPlayerRepository()

	service := command.NewRPCService(container)
	player := application.NewPlayerApplication(container.Hub(), playerRepo)
	controller := controller.NewWebSocketController(&service.RPC, container.Hub(), player)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			container.Engine().Stop()
			container.Hub().Stop()
			return nil
		},
	})

	return controller
}
