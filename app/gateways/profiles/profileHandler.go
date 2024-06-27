package httpProfile

import (
	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/profiles"
)

type ProfileHandler struct {
	service profiles.Service
}

// Account routes
func NewHandler(r chi.Router, usecase profiles.Service) *ProfileHandler {
	h := &ProfileHandler{
		service: usecase,
	}

	r.Post("/profile/create", h.CreateProfile)
	r.Get("/profile", h.GetProfile)
	r.Get("/profiles", h.GetAllProfiles)
	r.Get("/profiles/notIn", h.GetAllProfilesNotInGroup)
	r.Post("/profile/validate", h.ValidateProfile)

	return h
}
