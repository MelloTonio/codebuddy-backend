package accountUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/domain/errors"
)

func (accSvc *Services) GetAccount(id string) (account.Account, error) {
	accSvc.log.WithField("op", "GetAccount")

	accountF, err := accSvc.accountRepo.GetById(id)

	if err != nil {
		err := errors.ErrAccountNotFound
		accSvc.log.WithError(err).Error("account not found")
		return account.Account{}, err

	}

	return accountF, nil

}
