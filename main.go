package main

import (
	"fmt"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/command"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"golang.org/x/net/context"
)

func main() {
	fx.New(
		fx.Provide(
			rpc.NewRPC,
			NewHTTPServer,
			NewHub,
			NewEngine,
			NewStore,
			repository.NewPlayerRepository,
			repository.NewMatchRepository,
			service.NewBroadcastService,
			service.NewRecoveryService,
			service.NewGameLoopService,
			usecase.NewPlayer,
			usecase.NewMatch,
			command.NewRPCService,
		),
		fx.Invoke(func(*echo.Echo) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, rpc *rpc.RPC) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))
	e.GET("/ws", rpc.Serve)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go e.Start(":8080")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return e.Shutdown(ctx)
			},
		},
	)

	return e
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
	store := store.NewStore()

	store.CreateTable("players")
	store.CreateTable("matches")

	return store
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

func NewEventRouter(lc fx.Lifecycle) *message.Router {
	router, err := message.NewRouter(
		message.RouterConfig{},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to start event router: %s", err)
		return nil
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return router.Run(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return router.Close()
		},
	})

	return router
}
