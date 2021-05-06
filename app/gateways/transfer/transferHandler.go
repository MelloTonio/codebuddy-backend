package httpTransfer

import (
	"github.com/go-chi/chi"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
)

type TransferHandler struct {
	service transfer.Service
}

func NewHandler(r chi.Router, usecase transfer.Service) *TransferHandler {
	h := &TransferHandler{
		service: usecase,
	}

	r.Get("/", h.ShowAccountTransfers)
	r.Post("/{accountDestinationID}", h.Transfer)

	return h
}
