package httpAuth

import (
	"github.com/go-chi/chi/v5"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
)

type AuthHandler struct {
	service access.Service
}

// Auth routes
func NewHandler(r chi.Router, usecase access.Service) *AuthHandler {
	h := &AuthHandler{
		service: usecase,
	}

	r.Post("/accounts/login", h.Login)

	return h
}
