package repository

import (
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewInMemoryMatchRepository,
	wire.Bind(new(usecase.MatchRepository), new(*InMemoryMatchRepository)),
)
