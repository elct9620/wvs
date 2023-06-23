//go:build wireinject
// +build wireinject

package usecases

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewRoom,
)
