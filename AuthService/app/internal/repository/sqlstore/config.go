package sqlstore

import "fmt"

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pwd  string `yaml:"password"`
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "5432",
		Name: "Postgres",
		User: "Postgres",
		Pwd:  "Postgres",
	}
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", c.Host, c.Port, c.Name, c.User, c.Pwd)
}
