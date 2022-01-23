package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
)

// type key string

var (
	roles = []string{"admin", "brewery", "user"}
)

func (h *Handler) Middleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		accessToken := strings.Split(auth, "Bearer ")

		if len(accessToken) != 2 {
			h.error(w, r, http.StatusUnauthorized, errors.ErrUnauthorized)
			return
		}

		c, err := h.tokenManager.ParseToken(accessToken[1])
		if err != nil {
			h.error(w, r, http.StatusUnauthorized, errors.ErrUnauthorized)
			return
		}

		access := checkAudience(c.Audience)
		if !access {
			h.error(w, r, http.StatusNotAcceptable, errors.ErrNotAcceptable)
			return
		}

		id, err := strconv.Atoi(c.Subject)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		ctx := context.WithValue(r.Context(), domain.UID("id"), id)
		hf(w, r.WithContext(ctx))
	}
}

func checkAudience(role string) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}

	return false
}
