package server

type Config struct {
	HttpPort string `yaml:"port"`
	LogLevel string `yaml:"logLevel"`
}

func NewConfig() *Config {
	return &Config{
		HttpPort: "8080",
		LogLevel: "debug",
	}
}
