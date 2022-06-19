package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/pkg/command"
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/context"
)

func main() {
	container := container.NewContainer()
	playerRepo := container.NewPlayerRepository()

	service := command.NewRPCService(container)
	player := application.NewPlayerApplication(container.Hub(), playerRepo)
	controller := controller.NewWebSocketController(&service.RPC, container.Hub(), player)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))
	e.GET("/ws", controller.Server)

	var exit chan os.Signal = make(chan os.Signal, 1)
	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	signal.Notify(exit, os.Interrupt)
	<-exit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	container.Engine().Stop()
	container.Hub().Stop()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Error(err)
	}
}
