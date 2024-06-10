package challenges

import "context"

type Service interface {
	CreateChallenge(ctx context.Context, challenge Challenge) error
	SolveChallenge(ctx context.Context, challenge Challenge) (string, error)
	ListChallengesByGroup(ctx context.Context, groupName string, specificGroup string) ([]Challenge, error)
}
