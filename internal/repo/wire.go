//go:build wireinject
// +build wireinject

package repository

import "github.com/google/wire"

var ProvideInMemorySet = wire.NewSet(
	NewMemDB,
	NewInMemoryRoom,
	NewInMemoryPlayer,
)
