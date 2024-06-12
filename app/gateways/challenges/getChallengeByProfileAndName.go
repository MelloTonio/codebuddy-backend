package httpChallenge

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h ChallengeHandler) GetChallengesByProfileAndName(w http.ResponseWriter, r *http.Request) {
	groupName := r.URL.Query().Get("studyGroup")
	challengeName := r.URL.Query().Get("challengeName")
	profileName := r.URL.Query().Get("profileName")

	logrus.Infof("%s,%s,%s", groupName, challengeName, profileName)

	challenges, err := h.service.GetChallengesByGroupNameAndAlumni(context.Background(), groupName, challengeName, profileName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, challenges)
}
