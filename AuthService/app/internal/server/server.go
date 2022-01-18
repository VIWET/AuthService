package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	db     *sql.DB
	cache  *redis.Client
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

	if err := s.configureCache(); err != nil {
		return err
	}

	s.logger.Info(fmt.Sprintf("cache on %s:%s", s.config.CacheConfig.Host, s.config.CacheConfig.Port))

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

func (s *Server) configureCache() error {
	redis := redis.NewClient(&redis.Options{
		Addr:     s.config.CacheConfig.GetAddr(),
		Password: s.config.CacheConfig.Pwd,
		DB:       s.config.CacheConfig.DB,
	})

	status := redis.Ping()
	if err := status.Err(); err != nil {
		return err
	}

	s.cache = redis

	return nil
}
