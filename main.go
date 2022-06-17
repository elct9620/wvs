package main

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	container := container.NewContainer()
	playerRepo := container.NewPlayerRepository()

	game := application.NewGameApplication(container.Hub())
	match := application.NewMatchApplication(container.Hub())
	player := application.NewPlayerApplication(container.Hub(), playerRepo)
	controller := controller.NewWebSocketController(game, match, player)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))
	e.GET("/ws", controller.Server)
	e.Logger.Fatal(e.Start(":8080"))
}
