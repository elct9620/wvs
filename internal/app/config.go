package app

import (
	"github.com/elct9620/wvs/internal/config"
)

type Config struct {
	Address    string
	SessionKey string
}

func NewConfig(p config.Provider) *Config {
	return &Config{
		Address:    p.GetString(config.HttpAddr),
		SessionKey: p.GetString(config.SessionKey),
	}
}
