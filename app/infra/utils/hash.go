package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(secret string) (string, error) {
	password := []byte(secret)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)

	return string(hashedPassword), err
}
