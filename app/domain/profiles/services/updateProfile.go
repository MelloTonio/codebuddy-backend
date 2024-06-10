package services

import (
	"context"
)

func (sgs *ProfileService) UpdateProfileByUsername(ctx context.Context, username, groupName string) error {
	err := sgs.ProfileRepository.UpdateProfileByUsername(ctx, username, groupName)
	if err != nil {
		return err
	}

	return nil
}
