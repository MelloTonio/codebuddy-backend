package services

import (
	"context"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
)

type ChallengeService struct {
	ChallengeRepository challenges.Repository
}

func NewChallengeService(ChallengeRepository challenges.Repository) challenges.Service {
	return &ChallengeService{
		ChallengeRepository: ChallengeRepository,
	}
}

func (sgs *ChallengeService) CreateChallenge(ctx context.Context, challenge challenges.Challenge) error {
	err := sgs.ChallengeRepository.CreateChallenge(ctx, challenge)
	if err != nil {
		panic(err)
	}

	return nil
}

func (sgs *ChallengeService) ListChallengesByGroup(ctx context.Context, groupName string, specificGroup string) ([]challenges.Challenge, error) {
	dbChallenges, err := sgs.ChallengeRepository.ListChallengesByGroup(ctx, groupName)
	if err != nil {
		panic(err)
	}

	specificChallenge := []challenges.Challenge{}
	if specificGroup != "" {
		for _, challenge := range dbChallenges {
			if challenge.Name == specificGroup {
				specificChallenge = append(specificChallenge, challenge)
				return specificChallenge, nil
			}
		}
	}

	return dbChallenges, nil
}
