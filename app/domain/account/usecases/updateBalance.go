package accountUsecases

import "github.com/mellotonio/desafiogo/app/domain/account"

func (accSvc *Services) UpdateBalance(acc account.Account) error {
	accSvc.log.WithField("op", "ShowAccounts")

	err := accSvc.accountRepo.UpdateBalance(&acc)

	if err != nil {
		accSvc.log.WithError(err).Error("error while updating account balance")
		return err
	}

	return nil
}
