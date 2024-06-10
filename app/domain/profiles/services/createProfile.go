package services

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/mellotonio/desafiogo/app/domain/profiles"
)

type ProfileService struct {
	ProfileRepository profiles.Repository
}

func NewProfileService(ProfileRepository profiles.Repository) profiles.Service {
	return &ProfileService{ProfileRepository: ProfileRepository}
}

func (sgs *ProfileService) CreateProfile(ctx context.Context, profile profiles.Profile) error {
	hashedPassword, err := hashPassword(profile.Password)
	if err != nil {
		return err
	}

	profile.Password = hashedPassword

	err = sgs.ProfileRepository.CreateProfile(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

// hashPassword hashes a plain text password using bcrypt
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
