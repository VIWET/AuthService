package server

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Server) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info(fmt.Sprintf("database on: %s", s.config.DBConfig.GetConnectionString()))

	s.logger.Info(fmt.Sprintf("logger configured on level: %s", s.config.LogLevel))

	return nil
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
