package server

import "github.com/VIWET/Beeracle/AuthService/internal/repository/sqlstore"

type Config struct {
	HttpPort string          `yaml:"port"`
	LogLevel string          `yaml:"logLevel"`
	DBConfig sqlstore.Config `yaml:"db"`
}

func NewConfig() *Config {
	return &Config{
		HttpPort: "8080",
		LogLevel: "debug",
	}
}
