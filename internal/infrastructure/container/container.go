package container

import (
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
)

type Container struct {
	hub   *hub.Hub
	store *store.Store
}

func NewContainer() *Container {
	return &Container{
		hub:   hub.NewHub(),
		store: infrastructure.InitStore(),
	}
}

func (c *Container) Hub() *hub.Hub {
	return c.hub
}

func (c *Container) NewPlayerRepository() *repository.PlayerRepository {
	return repository.NewPlayerRepository(c.store)
}

func (c *Container) NewMatchRepository() *repository.MatchRepository {
	return repository.NewMatchRepository(c.store)
}
