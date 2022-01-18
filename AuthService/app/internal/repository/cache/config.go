package cache

import (
	"fmt"
	"time"
)

type Config struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	DB      int           `yaml:"db"`
	Pwd     string        `yaml:"password"`
	Expires time.Duration `yaml:"exp"`
}

func NewConfig() *Config {
	return &Config{
		Host:    "localhost",
		Port:    "6379",
		DB:      1,
		Pwd:     "",
		Expires: time.Duration(15) * time.Minute,
	}
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
