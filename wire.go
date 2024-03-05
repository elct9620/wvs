//go:build wireinject
// +build wireinject

package wvs

import (
	"github.com/elct9620/wvs/internal/app"
	"github.com/elct9620/wvs/internal/config"
	"github.com/google/wire"
)

func InitializeTest() (*app.Application, error) {
	wire.Build(
		config.DefaultSet,
		app.DefaultSet,
	)
	return nil, nil
}
