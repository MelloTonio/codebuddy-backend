package httpProfile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL)

	// Extract studentName from URL path
	username := r.URL.Query().Get("username")
	logrus.Infof("%s", username)

	profile, err := h.service.GetProfileByUsername(context.Background(), username)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.JSON(w, http.StatusCreated, profile)
}
