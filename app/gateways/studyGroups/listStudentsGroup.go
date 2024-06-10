package httpStudyGroup

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h StudyGroupHandler) ListStudentGroups(w http.ResponseWriter, r *http.Request) {
	// Extract studentName from URL path
	studentName := r.URL.Query().Get("studyGroup")

	logrus.Infof("%s", studentName)
	studyGroups, err := h.service.ListStudentGroups(context.Background(), studentName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	logrus.Infof("%+v", studyGroups)
	response.JSON(w, http.StatusCreated, studyGroups)
}
