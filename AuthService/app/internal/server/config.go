package server

import (
	"time"

	"github.com/VIWET/Beeracle/AuthService/internal/repository/cache"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/sqlstore"
)

type Config struct {
	HttpPort    string          `yaml:"port"`
	LogLevel    string          `yaml:"logLevel"`
	Salt        string          `yaml:"salt"`
	TokenExp    time.Duration   `yaml:"token_exp"`
	DBConfig    sqlstore.Config `yaml:"db"`
	CacheConfig cache.Config    `yaml:"cache"`
}

func NewConfig() *Config {
	return &Config{
		HttpPort: "8080",
		LogLevel: "debug",
	}
}
