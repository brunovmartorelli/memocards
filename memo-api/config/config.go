package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Mongo  Mongo  `mapstructure:",squash"`
	Server Server `mapstructure:",squash"`
}

type Server struct {
	Port string `mapstructure:"server_port"`
}

type Mongo struct {
	URI string `mapstructure:"mongo_uri"`
}

func New() (*Config, error) {
	viper.SetDefault("SERVER_PORT", "3000")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("config: .env file not found")
	}
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
