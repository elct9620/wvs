package config

import "github.com/google/wire"

var DefaultSet = wire.NewSet(
	New,
)

type Config struct {
	Http Http
}

func New() *Config {
	return &Config{
		Http: Http{
			Address: ":8080",
		},
	}
}
