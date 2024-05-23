package services

import (
	"context"

	studygroups "github.com/mellotonio/desafiogo/app/domain/studyGroups"
	"github.com/sirupsen/logrus"
)

type StudyGroupService struct {
	StudyGroupRepository studygroups.Repository
}

func NewStudyGroupService(StudyGroupRepository studygroups.Repository) studygroups.Service {
	return &StudyGroupService{
		StudyGroupRepository: StudyGroupRepository,
	}
}

func (sgs *StudyGroupService) SaveStudyGroup(ctx context.Context, studyGroup studygroups.StudyGroup) error {
	err := sgs.StudyGroupRepository.SaveStudyGroup(ctx, studyGroup)
	if err != nil {
		panic(err)
	}

	logrus.Infof(studyGroup.Name)

	return nil
}

func (sgs *StudyGroupService) ListStudentGroups(ctx context.Context, studentName string) ([]studygroups.StudyGroup, error) {
	studyGroups, err := sgs.StudyGroupRepository.GetStudyGroupsByStudent(ctx, studentName)
	if err != nil {
		panic(err)
	}

	return studyGroups, nil
}

func (sgs *StudyGroupService) ListStudentGroupDetails(ctx context.Context, studyGroupName string) (studygroups.StudyGroup, error) {
	studyGroups, err := sgs.StudyGroupRepository.GetStudyGroupDetails(ctx, studyGroupName)
	if err != nil {
		panic(err)
	}

	return studyGroups, nil
}

func (sgs *StudyGroupService) SaveWarning(ctx context.Context, studyGroupName, warnMessage string) error {
	err := sgs.StudyGroupRepository.SaveWarning(ctx, studyGroupName, warnMessage)
	if err != nil {
		panic(err)
	}

	return nil
}

func (sgs *StudyGroupService) GetWarnings(ctx context.Context, studyGroupName string) ([]string, error) {
	warnings, err := sgs.StudyGroupRepository.GetWarnings(ctx, studyGroupName)
	if err != nil {
		panic(err)
	}

	return warnings, nil
}
