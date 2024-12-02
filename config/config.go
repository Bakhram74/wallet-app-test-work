package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name    string `mapstructure:"APP_NAME"`
	Version string `mapstructure:"APP_VERSION"`
	Level   string `mapstructure:"LOG_LEVEL"`
	PGUrl     string `mapstructure:"PG_URL"`
	PoolMax int    `mapstructure:"PG_POOL_MAX"`
	Port    string `mapstructure:"HTTP_PORT"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
