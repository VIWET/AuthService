package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	db     *sql.DB
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

	s.logger.Info(fmt.Sprintf("logger configured on level: %s", s.config.LogLevel))

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info(fmt.Sprintf("database on %s:%s", s.config.DBConfig.Host, s.config.DBConfig.Port))

	s.logger.Info(fmt.Sprintf("Serving at http://localhost%s/", s.config.HttpPort))

	r := mux.NewRouter()

	return http.ListenAndServe(s.config.HttpPort, r)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureStore() error {
	db, err := sql.Open("postgres", s.config.DBConfig.GetConnectionString())
	if err != nil {
		s.logger.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		s.logger.Fatal(err)
	}

	s.db = db

	return nil
}
