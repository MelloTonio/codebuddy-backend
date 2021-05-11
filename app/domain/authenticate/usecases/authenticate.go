package usecasesAuth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/infra/utils"
)

func (as *accessService) Authenticate(cred access.Credential) (string, error) {

	log := as.log.WithField("op", "Authenticate").WithField("cpf", cred.CPF)

	ac, err := as.accountRepository.GetByCPF(cred.CPF)

	if err != nil {
		log.WithError(err).Error("There is no account related to this CPF")
		return "", err
	}

	if !utils.PasswordMatch(ac.Secret, cred.Secret) {
		err := errors.ErrPasswordsDontMatch
		log.WithError(err).Error("Passwords don't match")
		return "", err
	}

	atClaims := jwt.MapClaims{}
	atClaims["Cpf"] = ac.Cpf
	atClaims["Name"] = ac.Name
	atClaims["AccountID"] = ac.Id

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.WithError(err).Error("Error while signing jwt")
		return "", err
	}

	return token, nil
}
