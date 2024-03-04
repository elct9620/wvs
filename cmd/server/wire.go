//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/wvs/app"
	"github.com/google/wire"
)

func Initialize() (*app.Application, error) {
	wire.Build(app.DefaultSet)
	return nil, nil
}
