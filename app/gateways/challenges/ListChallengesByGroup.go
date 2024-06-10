package httpChallenge

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h ChallengeHandler) ListChallengesByGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL)
	groupName := r.URL.Query().Get("groupName")
	challengeName := r.URL.Query().Get("challengeName")

	studyGroups, err := h.service.ListChallengesByGroup(context.Background(), groupName, challengeName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	logrus.Infof("%+v", studyGroups)
	response.JSON(w, http.StatusCreated, studyGroups)
}
