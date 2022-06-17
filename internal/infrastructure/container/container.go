package container

import (
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
)

type Container struct {
	hub         *hub.Hub
	store       *store.Store
	playerStore *store.Table
	matchStore  *store.Table
}

func NewContainer() *Container {
	return &Container{
		hub:         hub.NewHub(),
		store:       infrastructure.InitStore(),
		playerStore: store.NewTable(),
		matchStore:  store.NewTable(),
	}
}

func (c *Container) Hub() *hub.Hub {
	return c.hub
}

func (c *Container) NewPlayerRepository() *repository.PlayerRepository {
	return repository.NewPlayerRepository(c.playerStore)
}

func (c *Container) NewMatchRepository() *repository.MatchRepository {
	return repository.NewMatchRepository(c.matchStore)
}
