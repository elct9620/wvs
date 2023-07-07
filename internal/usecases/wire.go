//go:build wireinject
// +build wireinject

package usecases

import (
	repository "github.com/elct9620/wvs/internal/repo"
	"github.com/google/wire"
)

var ProvideInMemoryUseCase = wire.NewSet(
	repository.ProvideInMemorySet,
	wire.Bind(new(RoomRepository), new(*repository.InMemoryRooms)),
	NewRoom,
)
