package accountUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/sirupsen/logrus"
)

// Helper to be used in tests
var (
	accountRepo     = mem.NewAccountRepository(logrus.New())
	accountServices = NewAccountService(accountRepo)

	fake_account = account.Account{
		Id:      "555-444-333",
		Cpf:     "452.394.019-38",
		Name:    "Doyle Von Frankestein",
		Balance: 1005,
		Secret:  "biggest secret in the world",
	}
)

// Define services - Account services depends on Repo
type Services struct {
	log         *logrus.Entry
	accountRepo account.Repository // (presentation) -> Services -> Repo -> Domain
}

// Generate new Account services
func NewAccountService(
	repoAccount account.Repository,
) account.Service {

	return &Services{
		accountRepo: repoAccount,
		log:         logrus.NewEntry(logrus.New()),
	}
}
