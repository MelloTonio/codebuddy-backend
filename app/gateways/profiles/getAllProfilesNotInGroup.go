package httpProfile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
)

func (h ProfileHandler) GetAllProfilesNotInGroup(w http.ResponseWriter, r *http.Request) {
	groupName := r.URL.Query().Get("groupName")
	fmt.Println(groupName)
	profiles, err := h.service.GetAllProfilesNotInGroup(context.Background(), groupName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.JSON(w, http.StatusCreated, profiles)
}
