package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
)

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// Extract token info
func ExtractTokenMetadata(r *http.Request) (access.Description, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return access.Description{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accID, ok := claims["AccountID"].(string)
		name, ok := claims["Name"].(string)
		cpf, ok := claims["Cpf"].(string)

		if !ok {
			return access.Description{}, err
		}

		return access.Description{
			AccountID: accID,
			Name:      name,
			CPF:       cpf,
		}, nil
	}

	return access.Description{}, err
}
