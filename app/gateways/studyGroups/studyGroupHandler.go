package httpStudyGroup

import (
	"github.com/go-chi/chi/v5"
	studygroups "github.com/mellotonio/desafiogo/app/domain/studyGroups"
)

type StudyGroupHandler struct {
	service studygroups.Service
}

// Account routes
func NewHandler(r chi.Router, usecase studygroups.Service) *StudyGroupHandler {

	h := &StudyGroupHandler{
		service: usecase,
	}

	r.Post("/studygroup/create", h.CreateStudyGroup)
	r.Get("/studygroup", h.ListStudents)
	r.Get("/studygroup/warnings/{studyGroup}", h.GetWarnings)
	r.Get("/studygroup/details", h.ListGroupDetails)
	r.Post("/studygroup/create/warning", h.CreateWarning)

	return h
}
