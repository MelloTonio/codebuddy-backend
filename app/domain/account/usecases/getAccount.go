package accountUsecases

import "github.com/mellotonio/desafiogo/app/domain/account"

func (accSvc *Services) GetAccount(id string) (account.Account, error) {
	accSvc.log.WithField("op", "GetAccount")

	account, err := accSvc.accountRepo.GetById(id)

	if err != nil {
		accSvc.log.WithError(err).Error("account not found")
	}

	return account, nil

}
