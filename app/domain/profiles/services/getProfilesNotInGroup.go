package services

import (
	"context"

	"github.com/mellotonio/desafiogo/app/domain/profiles"
)

func (sgs *ProfileService) GetAllProfilesNotInGroup(ctx context.Context, groupName string) ([]profiles.Profile, error) {
	storedProfiles, err := sgs.ProfileRepository.GetAllProfilesNotInGroup(ctx, groupName)
	if err != nil {
		return nil, err
	}

	return storedProfiles, nil
}
