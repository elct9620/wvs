package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var DefaultSet = wire.NewSet(
	NewViper,
	wire.Bind(new(Provider), new(*viper.Viper)),
)

type Provider interface {
	GetString(key string) string
}
