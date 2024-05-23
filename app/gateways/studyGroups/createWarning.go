package httpStudyGroup

import (
	"context"
	"net/http"

	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h StudyGroupHandler) CreateWarning(w http.ResponseWriter, r *http.Request) {
	var studyGroup StudyGroup

	err := response.Decode(r, &studyGroup)
	if err != nil {
		logrus.Infof("%s", err.Error())
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	logrus.Infof("%+v", studyGroup)
	err = h.service.SaveWarning(context.TODO(), studyGroup.Name, studyGroup.Warning)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, studyGroup)
}
