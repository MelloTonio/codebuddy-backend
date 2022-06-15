package httpCard

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/card"
)

type CardHandler struct {
	cardUC card.Usecase
}

// Account routes
func NewHandler(r chi.Router, usecase card.Usecase) *CardHandler {

	h := &CardHandler{
		cardUC: usecase,
	}

	r.Post("/cards/create", h.CreateCard)
	r.Get("/cards", h.GetAllCards)
	r.Get("/cards/tradeCard", h.TradeCard)
	r.Delete("/cards", h.DeleteCard)

	return h
}
