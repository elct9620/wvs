package container

import (
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
)

type Container struct {
	hub         *hub.Hub
	playerStore *store.Store
}

func NewContainer() *Container {
	return &Container{
		hub:         hub.NewHub(),
		playerStore: store.NewStore(),
	}
}

func (c *Container) Hub() *hub.Hub {
	return c.hub
}

func (c *Container) NewPlayerRepository() *repository.PlayerRepository {
	return repository.NewPlayerRepository(c.playerStore)
}
