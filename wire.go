//go:build wireinject
// +build wireinject

package wvs

import (
	"github.com/elct9620/wvs/app"
	"github.com/google/wire"
)

func InitializeTest() (*app.Application, error) {
	wire.Build(app.DefaultSet)
	return nil, nil
}
