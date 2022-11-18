package container

import (
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/service"
)

type Container struct {
	hub   *hub.Hub
	store *store.Store
}

func NewContainer(hub *hub.Hub, store *store.Store) *Container {
	return &Container{
		hub:   hub,
		store: store,
	}
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
