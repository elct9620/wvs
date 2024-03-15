package repository

import (
	"github.com/elct9620/wvs/internal/repository/inmemory"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	inmemory.NewMatchRepository,
	wire.Bind(new(usecase.MatchRepository), new(*inmemory.MatchRepository)),
)
