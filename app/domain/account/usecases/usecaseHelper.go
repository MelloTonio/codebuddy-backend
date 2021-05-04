package accountUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/sirupsen/logrus"
)

type Services struct {
	log         *logrus.Entry
	accountRepo account.Repository
}

func NewAccountService(
	repositoryRegistry account.Repository,
) account.Service {

	return &Services{
		accountRepo: repositoryRegistry,
		log:         logrus.NewEntry(&logrus.Logger{}),
	}
}
