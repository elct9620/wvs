package app

import "github.com/elct9620/wvs/internal/config"

const (
	ConfigHttpAddr = "http.addr"
)

type Config struct {
	Address string
}

func NewConfig(p config.Provider) *Config {
	return &Config{
		Address: p.GetString(ConfigHttpAddr),
	}
}
