package repository

import (
	"github.com/elct9620/wvs/internal/repository/inmemory"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	inmemory.DefaultSet,
)
