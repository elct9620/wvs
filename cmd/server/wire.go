//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/wvs/internal/app"
	"github.com/elct9620/wvs/internal/config"
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/eventbus"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/google/wire"
)

func Initialize() (*app.Application, error) {
	wire.Build(
		db.DefaultSet,
		eventbus.DefaultSet,
		repository.DefaultSet,
		usecase.DefaultSet,
		config.DefaultSet,
		app.DefaultSet,
	)
	return nil, nil
}
