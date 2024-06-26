package httpStudyGroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h StudyGroupHandler) ListGroupStudents(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL)

	// Extract studentName from URL path
	studentName := r.URL.Query().Get("studyGroup")

	studyGroups, err := h.service.ListGroupStudents(context.Background(), studentName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	logrus.Infof("%+v", studyGroups)
	response.JSON(w, http.StatusCreated, studyGroups)
}
