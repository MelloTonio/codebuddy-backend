package challenges

import "context"

type Repository interface {
	CreateChallenge(ctx context.Context, challenge Challenge) error
	ListChallengesByGroup(ctx context.Context, groupName string) ([]Challenge, error)
	UpdateChallenge(ctx context.Context, challenge Challenge) error
	GetChallengesByGroupNameAndAlumni(ctx context.Context, groupName, challengeName, alumniName string) ([]Challenge, error)
}
