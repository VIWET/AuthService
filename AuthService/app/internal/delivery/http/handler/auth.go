package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
)

func (h *Handler) SignUpUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &domain.UserCreateDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		_, t, err := h.services.User.SignUp(context.Background(), req, "user")
		if err != nil {
			// TODO: Error processing
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		// TODO: Custom responce for tokens
		h.respond(w, r, http.StatusOK, t)
	}
}
