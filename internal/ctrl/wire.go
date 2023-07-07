//go:build wireinject
// +build wireinject

package controller

import "github.com/google/wire"

var ProvideController = wire.NewSet(
	NewSystem,
	NewLobby,
)
