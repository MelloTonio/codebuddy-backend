package authenticationUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/sirupsen/logrus"
)

type accessService struct {
	log               *logrus.Entry
	accountRepository account.Repository
}

func NewAccessService(accountRepo account.Repository) access.Service {
	return &accessService{
		log:               logrus.NewEntry(logrus.New()),
		accountRepository: accountRepo,
	}
}

func (as *accessService) Authenticate(cred access.Credential) (access.Description, error) {
	var description access.Description

	log := as.log.WithField("op", "Authenticate").WithField("cpf", cred.CPF)

	// ToDo Validate credentials

	ac, err := as.accountRepository.GetByCPF(cred.CPF)

	if err != nil {
		log.WithError(err).Error("There is no account related to this CPF")
		return access.Description{}, err
	}

	if !utils.PasswordMatch(ac.Secret, cred.Secret) {
		err := errors.ErrPasswordsDontMatch
		log.WithError(err).Error("Passwords don't match")
		return access.Description{}, err
	}

	description.AccountID = ac.Id
	description.CPF = ac.Cpf
	description.Name = ac.Name

	return description, nil
}
