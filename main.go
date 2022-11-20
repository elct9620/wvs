package main

import (
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/server"
	"github.com/elct9620/wvs/internal/server/command"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/elct9620/wvs/pkg/store"
	"go.uber.org/fx"
	"golang.org/x/net/context"
)

func main() {
	fx.New(
		fx.Provide(
			NewHub,
			NewEngine,
			NewStore,
			repository.NewPlayerRepository,
			repository.NewMatchRepository,
			service.NewBroadcastService,
			service.NewRecoveryService,
			usecase.NewPlayer,
			usecase.NewMatch,
			AsRPCHandler(command.NewLoginCommand),
			AsRPCHandler(command.NewFindMatchCommand),
			AsRPCHandler(command.NewJoinMatchCommand),
			fx.Annotate(
				NewRPC,
				fx.ParamTags("hub", `group:"handlers"`),
			),
			NewHTTPServer,
		),
		fx.Invoke(func(*server.Server) {}),
	).Run()
}

func AsRPCHandler(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(rpc.CommandHandler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

func NewHTTPServer(lc fx.Lifecycle, rpc *rpc.RPC) *server.Server {
	server := server.NewServer(rpc)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.Start()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)

	return server
}

func NewRPC(hub *hub.Hub, handlers []rpc.CommandHandler) *rpc.RPC {
	rpc := rpc.NewRPC(hub)

	for _, handler := range handlers {
		rpc.Handle(handler)
	}

	return rpc
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
