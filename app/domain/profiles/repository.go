package profiles

import "context"

type Repository interface {
	CreateProfile(ctx context.Context, profile Profile) error
	GetProfile(ctx context.Context, profile Profile) (Profile, error)
	UpdateProfileByUsername(ctx context.Context, username string, groupName string) error
	GetAllProfiles(ctx context.Context) ([]Profile, error)
}
