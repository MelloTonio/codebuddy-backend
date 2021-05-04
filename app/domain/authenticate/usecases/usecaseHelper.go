package authenticationUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/sirupsen/logrus"
)

type accessService struct {
	log               *logrus.Entry
	accountRepository account.Repository
}

// NewAccessService returns a new access service.
func NewAccessService(accountRepo account.Repository) access.Service {
	return &accessService{
		log:               logrus.NewEntry(&logrus.Logger{}),
		accountRepository: accountRepo,
	}
}

func (as *accessService) Authenticate(cred access.Credential) (d access.Description, err error) {
	log := as.log.WithField("op", "Authenticate").WithField("cpf", cred.CPF)

	// Validate the struct cred

	if err != nil {
		return
	}

	// ac, err := as.accountRepository.GetByCPF(cred.CPF)
	if err != nil {
		log.WithError(err).Error("no account found with this cpf")
		return
	}

	// Match hash and secret
	return
}
