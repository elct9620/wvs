//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/wvs/internal/app"
	"github.com/elct9620/wvs/internal/config"
	"github.com/google/wire"
)

func Initialize() (*app.Application, error) {
	wire.Build(
		config.DefaultSet,
		app.DefaultSet,
	)
	return nil, nil
}
