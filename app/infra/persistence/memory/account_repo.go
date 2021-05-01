package mem

import (
	"time"

	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/sirupsen/logrus"
)

type memAccountRepo struct {
	accounts []account.Account
	log      *logrus.Entry
}

func NewAccountRepository(logger *logrus.Logger) account.Repository {
	return &memAccountRepo{
		accounts: []account.Account{},
		log:      logger.WithField("source", "memAccountRepository"),
	}
}

func (memoryRepo *memAccountRepo) Store(_account *account.Account) error {
	if _account.Id == "" {
		err := errors.EmptyAccountID_Err
		memoryRepo.log.WithError(err).Error("Empty Account Id")
		return err
	}

	index := memoryRepo.findByCpf(_account.Cpf)

	if index != -1 {
		err := errors.AccountAlreadyExists_Err
		memoryRepo.log.WithError(err).Error("Account already exists")
		return err
	}

	_account.Created_at = time.Now()

	memoryRepo.accounts = append(memoryRepo.accounts, *_account)

	//memoryRepo.log.Info("Account %s created sucessfully", account.Id)

	return nil
}

func (memoryRepo *memAccountRepo) GetById(id string) (account.Account, error) {
	if id == "" {
		err := errors.EmptyAccountID_Err
		memoryRepo.log.WithError(err).Error("Empty Account Id")
		return account.Account{}, err
	}

	index := memoryRepo.find(id)

	if index == -1 {
		err := errors.AccountNotFound_Err
		memoryRepo.log.WithError(err).Error("Account not found")
		return account.Account{}, err
	}

	return memoryRepo.accounts[index], nil

}

func (memoryRepo *memAccountRepo) ExistsByCPF(_account *account.Account) (bool, error) {
	if _account.Id == "" {
		err := errors.EmptyAccountID_Err
		memoryRepo.log.WithError(err).Error("Empty Account Id")
		return true, err
	}

	index := memoryRepo.findByCpf(_account.Cpf)

	// If the index is different than -1, the account already exists
	if index != -1 {
		err := errors.AccountAlreadyExists_Err
		memoryRepo.log.WithError(err).Error("Account already exists")
		return true, nil
	}

	return false, nil

}

func (memoryRepo *memAccountRepo) GetByCPF(cpf string) (account.Account, error) {
	if cpf == "" {
		err := errors.EmptyCPF_Err
		memoryRepo.log.WithError(err).Error("Empty CPF value")
		return account.Account{}, err
	}

	index := memoryRepo.findByCpf(cpf)

	if index == -1 {
		err := errors.AccountNotFound_Err
		memoryRepo.log.WithError(err).Error("Account Not Found")
		return account.Account{}, err
	}

	return memoryRepo.accounts[index], nil

}

func (memoryRepo *memAccountRepo) UpdateBalance(_account *account.Account, balance int) error {
	if _account.Id == "" {
		err := errors.EmptyAccountID_Err
		memoryRepo.log.WithError(err).Error("Empty Account Id")
		return err
	}

	index := memoryRepo.find(_account.Id)

	if index == -1 {
		err := errors.AccountNotFound_Err
		memoryRepo.log.WithError(err).Error("Account Not Found")
	}

	memoryRepo.accounts[index].Balance = balance

	return nil

}

func (memoryRepo *memAccountRepo) GetBalance(account account.Account) (int, error) {
	if account.Id == "" {
		err := errors.EmptyAccountID_Err
		memoryRepo.log.WithError(err).Error("Empty Account Id")
		return -1, err
	}

	index := memoryRepo.find(account.Id)

	if index == -1 {
		err := errors.AccountNotFound_Err
		memoryRepo.log.WithError(err).Error("Account Not Found")
		return -1, err
	}

	return memoryRepo.accounts[index].Balance, nil
}

func (mar *memAccountRepo) GenerateID() string {
	return utils.GenUUID()
}

func (memoryRepo *memAccountRepo) ShowAll() ([]account.Account, error) {
	accounts := make([]account.Account, len(memoryRepo.accounts))

	copy(accounts, memoryRepo.accounts)

	return accounts, nil
}

func (memoryRepo *memAccountRepo) find(accountID string) int {
	for index, account := range memoryRepo.accounts {
		if account.Id == accountID {
			return index
		}
	}

	return -1
}

func (memoryRepo *memAccountRepo) findByCpf(accountCPF string) int {
	for index, account := range memoryRepo.accounts {
		if account.Cpf == accountCPF {
			return index
		}
	}

	return -1
}
