package services

import (
	"context"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
	"github.com/sirupsen/logrus"
)

func (sgs *ChallengeService) UpdateChallenge(ctx context.Context, challenge challenges.Challenge) error {
	err := sgs.ChallengeRepository.UpdateChallenge(ctx, challenge)
	if err != nil {
		logrus.Infof("%+s", err.Error())
		return nil
	}

	return nil
}
