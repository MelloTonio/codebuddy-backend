package services

import (
	"context"

	"github.com/mellotonio/desafiogo/app/domain/profiles"
)

func (sgs *ProfileService) GetProfileByUsername(ctx context.Context, username string) (profiles.Profile, error) {
	storedProfile, err := sgs.ProfileRepository.GetProfile(ctx, profiles.Profile{
		Username: username,
	})
	if err != nil {
		return profiles.Profile{}, err
	}

	return storedProfile, nil
}
