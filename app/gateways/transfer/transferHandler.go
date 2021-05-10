package httpTransfer

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
)

type TransferHandler struct {
	service transfer.Service
}

func NewHandler(r chi.Router, usecase transfer.Service) *TransferHandler {
	h := &TransferHandler{
		service: usecase,
	}

	r.Get("/transfers/{accountID}", h.ShowAccountTransfers)
	r.Post("/transferTo/{accountDestinationID}", h.Transfer)

	return h
}
