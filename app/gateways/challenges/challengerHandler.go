package httpChallenge

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/challenges"
)

type ChallengeHandler struct {
	service challenges.Service
}

// Account routes
func NewHandler(r chi.Router, usecase challenges.Service) *ChallengeHandler {

	h := &ChallengeHandler{
		service: usecase,
	}

	r.Post("/challenge/solve", h.SolveChallenge)
	r.Post("/challenge/create", h.CreateChallenge)
	r.Get("/challenges", h.ListChallengesByGroup)
	r.Get("/challenges", h.ListChallengesByGroup)

	return h
}
