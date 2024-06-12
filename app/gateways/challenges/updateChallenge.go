package httpChallenge

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h ChallengeHandler) UpdateChallenge(w http.ResponseWriter, r *http.Request) {
	var challengeBody Challenge

	err := response.Decode(r, &challengeBody)
	if err != nil {
		logrus.Infof("%s", err.Error())
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	logrus.Infof("%+v", challengeBody)
	err = h.service.UpdateChallenge(context.Background(), challenges.Challenge{
		Name:  challengeBody.Name,
		Group: challengeBody.Group,
		Answer: []challenges.Answer{
			{
				AlumniName: challengeBody.Answer[0].AlumniName,
				Text:       challengeBody.Answer[0].Text,
			},
		},
	})
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, challengeBody)
}
