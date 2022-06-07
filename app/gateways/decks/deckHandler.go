package httpDeck

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/deck"
)

type DeckHandler struct {
	deckUC deck.Usecase
}

// Account routes
func NewHandler(r chi.Router, usecase deck.Usecase) *DeckHandler {

	h := &DeckHandler{
		deckUC: usecase,
	}

	r.Post("/decks/create", h.CreateDeck)

	return h
}
