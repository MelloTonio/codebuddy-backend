package studygroups

import "context"

type Repository interface {
	SaveStudyGroup(ctx context.Context, stdGroup StudyGroup) error
	SaveWarning(ctx context.Context, studyGroupName, warnMessage string) error
	GetWarnings(ctx context.Context, studyGroupName string) ([]string, error)
	GetStudyGroupsByStudent(ctx context.Context, studentName string) ([]StudyGroup, error)
	GetStudyGroupDetails(ctx context.Context, studyGroupName string) (StudyGroup, error)
}
