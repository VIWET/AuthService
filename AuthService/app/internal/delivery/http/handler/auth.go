package handler

import (
	"encoding/json"
	"net/http"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
)

func (h *Handler) SignUpUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &domain.UserCreateDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		t, err := h.services.User.SignUp(r.Context(), req, "user", r.UserAgent())
		if err != nil {
			switch err {
			case errors.ErrEmailIsEmpty, errors.ErrEmailIsNotValid, errors.ErrPasswordIsEmpty, errors.ErrPasswordConfirmation:
				h.error(w, r, http.StatusBadRequest, err)
				return
			default:
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		c := http.Cookie{
			Name:     "RefreshToken",
			Value:    t.RefreshToken,
			Path:     "api/auth",
			MaxAge:   int(t.Exp),
			HttpOnly: true,
			// Secure:   true,
		}

		http.SetCookie(w, &c)
		h.respond(w, r, http.StatusOK, t)
	}
}

func (h *Handler) SignUpBrewery() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &domain.UserCreateDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		t, err := h.services.User.SignUp(r.Context(), req, "brewery", r.UserAgent())
		if err != nil {
			switch err {
			case errors.ErrEmailIsEmpty, errors.ErrEmailIsNotValid, errors.ErrPasswordIsEmpty, errors.ErrPasswordLength, errors.ErrPasswordConfirmation:
				h.error(w, r, http.StatusBadRequest, err)
				return
			default:
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		c := http.Cookie{
			Name:     "RefreshToken",
			Value:    t.RefreshToken,
			Path:     "api/auth",
			MaxAge:   int(t.Exp),
			HttpOnly: true,
			// Secure:   true,
		}

		http.SetCookie(w, &c)
		h.respond(w, r, http.StatusOK, t)
	}
}

func (h *Handler) Refresh() http.HandlerFunc {

	type Refresh struct {
		Fingerprint string `json:"fingerprint"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Refresh{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		cookie, err := r.Cookie("RefreshToken")
		if err != nil {
			h.error(w, r, http.StatusUnauthorized, err)
			return
		}

		t, err := h.services.User.Refresh(r.Context(), cookie.Value, r.UserAgent(), req.Fingerprint)
		if err != nil {
			switch err {
			case errors.ErrUnauthorized:
				h.error(w, r, http.StatusUnauthorized, err)
				return
			default:
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		c := http.Cookie{
			Name:     "RefreshToken",
			Value:    t.RefreshToken,
			Path:     "api/auth",
			MaxAge:   int(t.Exp),
			HttpOnly: true,
			// Secure:   true,
		}

		http.SetCookie(w, &c)
		h.respond(w, r, http.StatusOK, t)
	}
}

func (h *Handler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &domain.UserSignIn{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		t, err := h.services.User.SignIn(r.Context(), req, r.UserAgent())
		if err != nil {
			switch err {
			case errors.ErrUnauthorized:
				h.error(w, r, http.StatusUnauthorized, err)
				return
			default:
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		c := http.Cookie{
			Name:     "RefreshToken",
			Value:    t.RefreshToken,
			Path:     "api/auth",
			MaxAge:   int(t.Exp),
			HttpOnly: true,
			// Secure:   true,
		}

		http.SetCookie(w, &c)
		h.respond(w, r, http.StatusOK, t)
	}
}

func (h *Handler) Delete() http.HandlerFunc {

	type Delete struct {
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Delete{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		cookie, err := r.Cookie("AccessToken")
		if err != nil {
			h.error(w, r, http.StatusUnauthorized, err)
			return
		}

		if err := h.services.User.Delete(r.Context(), req.Password, cookie.Value); err != nil {
			switch err {
			case errors.ErrUnauthorized:
				h.error(w, r, http.StatusUnauthorized, err)
				return
			default:
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &domain.UserUpdateDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		cookie, err := r.Cookie("AccessToken")
		if err != nil {
			h.error(w, r, http.StatusUnauthorized, err)
			return
		}

		if err := h.services.User.Update(r.Context(), req, cookie.Value); err != nil {
			switch err {
			case errors.ErrUnauthorized:
				h.error(w, r, http.StatusUnauthorized, err)
				return
			default:
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
		}
	}
}
