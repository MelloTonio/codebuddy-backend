package accountUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/infra/utils"
)

func (accSvc *Services) CreateAccount(acc *account.Account) error {
	var err error
	acc.Id = accSvc.accountRepo.GenerateID()

	accSvc.log.WithField("op", "CreateAccount")

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
