package services

import (
	"context"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
)

func (sgs *ChallengeService) GetChallengesByGroupNameAndAlumni(ctx context.Context, groupName, challengeName, alumniName string) ([]challenges.Challenge, error) {
	challenges, err := sgs.ChallengeRepository.GetChallengesByGroupNameAndAlumni(ctx, groupName, challengeName, alumniName)
	if err != nil {
		return nil, err
	}

	return challenges, nil
}
