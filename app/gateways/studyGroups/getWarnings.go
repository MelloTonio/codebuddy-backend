package httpStudyGroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h StudyGroupHandler) GetWarnings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL)

	// Extract studentName from URL path
	studyGroupName := chi.URLParam(r, "studyGroup")

	warnings, err := h.service.GetWarnings(context.Background(), studyGroupName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	warningResponse := []respWarning{}
	for _, warningText := range warnings {
		warningResponse = append(warningResponse, respWarning{
			WarningText: warningText,
		})
	}

	response.JSON(w, http.StatusCreated, warningResponse)
}

type respWarning struct {
	WarningText string `json:"warning_text"`
}
