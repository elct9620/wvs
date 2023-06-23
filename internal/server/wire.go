//go:build wireinject
// +build wireinject

package server

import "github.com/google/wire"

var ProvideInMemorySession = wire.NewSet(
	NewInMemorySession,
	wire.Bind(new(SessionStore), new(*InMemorySession)),
)
