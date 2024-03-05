package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var DefaultSet = wire.NewSet(
	NewViperWithDefaults,
	wire.Bind(new(Provider), new(*viper.Viper)),
)

const (
	HttpAddr   = "http.addr"
	SessionKey = "session.key"
)

type Provider interface {
	GetString(key string) string
}
