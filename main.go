package main

import (
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	controller := controller.NewWebSocketController()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))
	e.GET("/ws", controller.Server)
	e.Logger.Fatal(e.Start(":8080"))
}
