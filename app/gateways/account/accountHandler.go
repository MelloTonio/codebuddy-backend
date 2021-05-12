package httpAccount

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/account"
)

type AccountHandler struct {
	service account.Service
}

// Account routes
func NewHandler(r chi.Router, usecase account.Service) *AccountHandler {

	h := &AccountHandler{
		service: usecase,
	}

	r.Post("/accounts/create", h.CreateAccount)
	r.Get("/accounts/{accountID}", h.GetAccount)
	r.Get("/accounts/all", h.ShowAccounts)

	return h
}
