package usecasesAcc

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/infra/utils"
)

func (accSvc *Services) CreateAccount(acc *account.Account) error {
	var err error

	// Condition to test "GetAccount"
	if len(acc.Id) == 0 {
		acc.Id = accSvc.accountRepo.GenerateID()
	}

	accSvc.log.WithField("op", "CreateAccount")

	// ToDo: VALIDATION ON ACCOUNT STRUCT
	if len(acc.Secret) < 5 {
		err := errors.ErrEmptyAccountSecret
		accSvc.log.WithError(err).WithField("Account Secret", acc.Id).Error("Account secret cannot be empty")
		return err
	}

	acc.Secret, err = utils.Hash(acc.Secret)

	if err != nil {
		accSvc.log.WithError(err).WithField("accountID", acc.Id).Error("Error hashing secret")
		return err
	}

	err = accSvc.accountRepo.Store(acc)

	if err != nil {
		accSvc.log.WithError(err).WithField("accountID", acc.Id).Error("Error in accountUsecases.services.Store")
		return err
	}

	return nil
}
