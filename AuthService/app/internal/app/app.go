package app

import (
	"flag"
	"io/ioutil"

	"github.com/VIWET/Beeracle/AuthService/internal/server"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to server config file")
}

func Run() {
	flag.Parse()

	config := server.NewConfig()

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		logrus.Fatal(err)
	}

	s := server.New(config)

	if err := s.Run(); err != nil {
		logrus.Fatal(err)
	}
}
