package httpStudyGroup

import (
	"context"
	"net/http"
	"strings"

	studygroups "github.com/mellotonio/desafiogo/app/domain/studyGroups"
	"github.com/mellotonio/desafiogo/app/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h StudyGroupHandler) CreateStudyGroup(w http.ResponseWriter, r *http.Request) {
	var accountBody StudyGroup

	err := response.Decode(r, &accountBody)
	if err != nil {
		logrus.Infof("%s", err.Error())
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	logrus.Infof("%+v", accountBody)
	err = h.service.SaveStudyGroup(context.Background(), studygroups.StudyGroup{
		Name:        accountBody.Name,
		Students:    splitStudents(accountBody.Students),
		Subject:     accountBody.Subject,
		Description: accountBody.Description,
	})
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, accountBody)
}

func splitStudents(students string) []string {
	// Split the string by commas and trim any leading/trailing spaces from each name
	split := strings.Split(students, ",")
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
	}
	return split
}
