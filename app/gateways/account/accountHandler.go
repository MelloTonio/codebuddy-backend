package httpAccount

import (
	"github.com/go-chi/chi"
	"github.com/mellotonio/desafiogo/app/domain/account"
)

type AccountHandler struct {
	service account.Service
}

func NewHandler(r *chi.Mux, usecase account.Service) *AccountHandler {

	h := &AccountHandler{
		service: usecase,
	}

	r.Post("/", h.CreateAccount)
	r.Get("/{accountID}", h.GetAccount)
	r.Get("/all", h.ShowAccounts)

	return h
}
