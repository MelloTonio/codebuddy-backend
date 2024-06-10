package httpChallenge

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h ChallengeHandler) CreateChallenge(w http.ResponseWriter, r *http.Request) {
	var challengeBody Challenge

	err := response.Decode(r, &challengeBody)
	if err != nil {
		logrus.Infof("%s", err.Error())
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	logrus.Infof("%+v", challengeBody)
	err = h.service.CreateChallenge(context.Background(), challenges.Challenge{
		Name:      challengeBody.Name,
		Text:      challengeBody.Text,
		Input:     challengeBody.Input,
		Output:    challengeBody.Output,
		Group:     challengeBody.Group,
		Difficult: challengeBody.Difficult,
	})
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, challengeBody)
}
