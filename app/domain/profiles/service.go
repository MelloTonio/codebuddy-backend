package profiles

import "context"

type Service interface {
	CreateProfile(ctx context.Context, challenge Profile) error
	GetProfileByUsername(ctx context.Context, username string) (Profile, error)
	ValidateProfile(ctx context.Context, profile Profile) (string, error)
	UpdateProfileByUsername(ctx context.Context, usernames string, groupName string) error
	GetAllProfiles(ctx context.Context) ([]Profile, error)
	GetAllProfilesNotInGroup(ctx context.Context, groupName string) ([]Profile, error)
}
