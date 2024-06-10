package services

import (
	"context"
	"errors"

	"github.com/mellotonio/desafiogo/app/domain/profiles"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (sgs *ProfileService) ValidateProfile(ctx context.Context, profile profiles.Profile) error {
	storedProfile, err := sgs.ProfileRepository.GetProfile(ctx, profile)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedProfile.Password), []byte(profile.Password))
	if err != nil {
		logrus.Errorf("password mismatch: %v", err)
		return errors.New("invalid credentials")
	}

	logrus.Infof("successfully validated profile: %v", profile.Username)
	return nil
}
