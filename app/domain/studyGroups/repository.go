package studygroups

import "context"

type Repository interface {
	SaveStudyGroup(ctx context.Context, stdGroup StudyGroup) error
	SaveWarning(ctx context.Context, studyGroupName, warnMessage string) error
	GetWarnings(ctx context.Context, studyGroupName string) ([]string, error)
	ListGroupStudents(ctx context.Context, groupName string) ([]StudyGroup, error)
	GetStudyGroupsByStudent(ctx context.Context, studentName string) ([]StudyGroup, error)
	AddStudentsToGroup(ctx context.Context, studyGroupName string, students []string) error
	GetStudyGroupDetails(ctx context.Context, studyGroupName string) (StudyGroup, error)
}
