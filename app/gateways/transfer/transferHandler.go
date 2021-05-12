package httpTransfer

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
)

type TransferHandler struct {
	service transfer.Service
}

// Transfer routes
func NewHandler(r chi.Router, usecase transfer.Service) *TransferHandler {
	h := &TransferHandler{
		service: usecase,
	}

	r.Get("/transfers/{accountID}", h.ShowAccountTransfers)

	// Example: /transferTo?accountID_destination=xyz&amount=50
	r.Post("/transferTo", h.Transfer)

	return h
}
