package httpProfile

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/domain/profiles"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profileBody Profile

	err := response.Decode(r, &profileBody)
	if err != nil {
		logrus.Infof("%s", err.Error())
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	logrus.Infof("%+v", profileBody)
	err = h.service.CreateProfile(context.Background(), profiles.Profile{
		Username:    profileBody.Username,
		Password:    profileBody.Password,
		Groups:      []string{},
		ProfileType: profileBody.ProfileType,
	})
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, profileBody)
}
