package main

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	playerRepo := repository.NewPlayerRepository()

	game := application.NewGameApplication()
	player := application.NewPlayerApplication(playerRepo)
	broadcast := application.NewBroadcastApplication(playerRepo)
	controller := controller.NewWebSocketController(game, player, broadcast)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))
	e.GET("/ws", controller.Server)
	e.Logger.Fatal(e.Start(":8080"))
}
