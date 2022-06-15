package httpUser

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/user"
)

type UserHandler struct {
	userUC user.Usecase
}

// user routes
func NewHandler(r chi.Router, usecase user.Usecase) *UserHandler {

	h := &UserHandler{
		userUC: usecase,
	}

	r.Get("/users/all", h.GetAllUsers)
	r.Get("/users", h.Get)

	return h
}
