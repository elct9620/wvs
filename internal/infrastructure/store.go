package infrastructure

import "github.com/elct9620/wvs/internal/infrastructure/store"

func InitStore() *store.Store {
	store := store.NewStore()

	store.CreateTable("players")
	store.CreateTable("matches")

	return store
}
