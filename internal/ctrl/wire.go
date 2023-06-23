//go:build wireinject
// +build wireinject

package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewSystem,
	NewLobby,
)
