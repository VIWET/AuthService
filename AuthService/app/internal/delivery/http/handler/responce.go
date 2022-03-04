package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	h.logger.Error(err)
	json.NewEncoder(w).Encode(&errorResponse{
		StatusCode: code,
		Message:    err.Error(),
	})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
