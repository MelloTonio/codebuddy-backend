package accountUsecases

import "github.com/mellotonio/desafiogo/app/domain/account"

func (accSvc *Services) ShowAccounts() ([]account.Account, error) {
	accSvc.log.WithField("op", "ShowAccounts")

	accounts, err := accSvc.accountRepo.ShowAll()

	if err != nil {
		accSvc.log.WithError(err).Error("cannot get accounts")
	}

	return accounts, nil
}
