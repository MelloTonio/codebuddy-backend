package studygroups

import "context"

type Service interface {
	SaveStudyGroup(ctx context.Context, studyGroup StudyGroup) error
	SaveWarning(ctx context.Context, studyGroupName, warnMessage string) error
	GetWarnings(ctx context.Context, studyGroupName string) ([]string, error)
	ListStudentGroups(ctx context.Context, studentName string) ([]StudyGroup, error)
	ListGroupStudents(ctx context.Context, groupName string) ([]StudyGroup, error)
	ListStudentGroupDetails(ctx context.Context, studyGroup string) (StudyGroup, error)
}
