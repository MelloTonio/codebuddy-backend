package httpProfile

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h ProfileHandler) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := h.service.GetAllProfiles(context.Background())
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.JSON(w, http.StatusCreated, profiles)
}
