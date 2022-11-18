package container

import (
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/service"
)

type Container struct {
	hub    *hub.Hub
	engine *engine.Engine
	store  *store.Store
}

func NewContainer(hub *hub.Hub, engine *engine.Engine, store *store.Store) *Container {
	return &Container{
		hub:    hub,
		engine: engine,
		store:  store,
	}
}

func (c *Container) Hub() *hub.Hub {
	return c.hub
}

func (c *Container) Engine() *engine.Engine {
	return c.engine
}

func (c *Container) NewBroadcastService() *service.BroadcastService {
	return service.NewBroadcastService(c.hub)
}

func (c *Container) NewGameLoopService() *service.GameLoopService {
	return service.NewGameLoopService(c.NewBroadcastService(), c.NewRecoveryService())
}

func (c *Container) NewRecoveryService() *service.RecoveryService {
	return service.NewRecoveryService(c.NewBroadcastService())
}
