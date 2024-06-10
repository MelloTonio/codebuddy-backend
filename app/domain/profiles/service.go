package profiles

import "context"

type Service interface {
	CreateProfile(ctx context.Context, challenge Profile) error
	GetProfileByUsername(ctx context.Context, username string) (Profile, error)
	ValidateProfile(ctx context.Context, profile Profile) error
	UpdateProfileByUsername(ctx context.Context, usernames string, groupName string) error
}
