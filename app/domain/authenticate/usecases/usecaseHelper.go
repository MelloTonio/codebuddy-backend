package authUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
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
