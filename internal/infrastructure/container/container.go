package container

import (
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
)

type Container struct {
	hub    *hub.Hub
	engine *engine.Engine
	store  *store.Store
}

func NewContainer() *Container {
	return &Container{
		hub:    hub.NewHub(),
		engine: engine.NewEngine(),
		store:  infrastructure.InitStore(),
	}
}

func (c *Container) Hub() *hub.Hub {
	return c.hub
}

func (c *Container) Engine() *engine.Engine {
	return c.engine
}

func (c *Container) NewPlayerRepository() *repository.PlayerRepository {
	return repository.NewPlayerRepository(c.store)
}

func (c *Container) NewMatchRepository() *repository.MatchRepository {
	return repository.NewMatchRepository(c.store)
}

func (c *Container) NewBroadcastService() *service.BroadcastService {
	return service.NewBroadcastService(c.hub)
}
