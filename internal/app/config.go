package app

import "github.com/elct9620/wvs/internal/config"

type Config struct {
	Address string
}

func NewConfig(config *config.Config) *Config {
	return &Config{
		Address: config.Http.Address,
	}
}
