package config

import "github.com/spf13/viper"

func NewViper() (*viper.Viper, error) {
	v := viper.New()

	v.AutomaticEnv()
	v.SetConfigName("server")
	v.SetConfigType("toml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	return v, nil
}
