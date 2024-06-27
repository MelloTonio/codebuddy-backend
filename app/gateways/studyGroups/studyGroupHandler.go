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
	r.Post("/studygroup/update", h.AddStudentsToGroup)
	r.Get("/studygroup/students", h.ListStudentGroups)
	r.Get("/studygroup/groups", h.ListGroupStudents)
	r.Get("/studygroup/warnings/{studyGroup}", h.GetWarnings)
	r.Get("/studygroup/details", h.ListGroupDetails)
	r.Post("/studygroup/create/warning", h.CreateWarning)

	return h
}
