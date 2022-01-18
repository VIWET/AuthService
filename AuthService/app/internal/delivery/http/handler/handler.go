package handler

import (
	"github.com/VIWET/Beeracle/AuthService/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	router   *mux.Router
	services *service.Services
	logger   *logrus.Logger
}

func New(services *service.Services, logger *logrus.Logger) *Handler {
	return &Handler{
		router:   mux.NewRouter().StrictSlash(true),
		services: services,
		logger:   logger,
	}
}

func (h *Handler) configureRouter() {
	auth := h.router.PathPrefix("/api/auth").Subrouter()

	auth.Handle("/user/sign-up", h.SignUpUser()).Methods("POST")
}

func (h *Handler) GetRouter() *mux.Router {
	h.configureRouter()
	return h.router
}
