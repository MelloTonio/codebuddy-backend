package httpStudyGroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h StudyGroupHandler) ListStudents(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL)

	// Extract studentName from URL path
	studentName := r.URL.Query().Get("studyGroup")

	// Now you can use studentName as needed, for example, pass it to your service method
	studyGroups, err := h.service.ListStudentGroups(context.Background(), studentName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response.JSON(w, http.StatusCreated, studyGroups)
}
