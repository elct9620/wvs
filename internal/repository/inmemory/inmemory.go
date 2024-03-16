package inmemory

import (
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewMatchRepository,
	wire.Bind(new(usecase.MatchRepository), new(*MatchRepository)),
)
