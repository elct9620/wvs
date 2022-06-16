package container

import (
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
)

type Container struct {
	playerStore *store.Store
}

func NewContainer() *Container {
	return &Container{
		playerStore: store.NewStore(),
	}
}

func (c *Container) GetPlayerRepository() *repository.PlayerRepository {
	return repository.NewPlayerRepository(c.playerStore)
}
