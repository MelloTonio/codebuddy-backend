package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mellotonio/desafiogo/app/domain/profiles"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "beniciogay"

func createToken(username, profile_type string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":     username,
			"profile_type": profile_type,
			"exp":          time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (sgs *ProfileService) ValidateProfile(ctx context.Context, profile profiles.Profile) (string, error) {
	storedProfile, err := sgs.ProfileRepository.GetProfile(ctx, profile)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedProfile.Password), []byte(profile.Password))
	if err != nil {
		logrus.Errorf("password mismatch: %v", err)
		return "", errors.New("invalid credentials")
	}

	jwtToken, err := createToken(storedProfile.Username, storedProfile.ProfileType)
	if err != nil {
		logrus.Errorf("error creating token jwt: %v", err)
		return "", err
	}

	logrus.Infof("successfully validated profile: %s", jwtToken)
	return jwtToken, nil
}
