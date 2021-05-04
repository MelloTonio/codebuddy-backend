package transferUsecases

import "github.com/mellotonio/desafiogo/app/domain/transfer"

func (transfSvc *Services) ShowAllAccountTransfers(Account_Id string) ([]transfer.Transfer, error) {

	transfSvc.log.WithField("op", "ShowAllAccountTransfers")

	transfers, err := transfSvc.transfRepo.ListByAccId(Account_Id)

	if err != nil {
		transfSvc.log.WithError(err).Error("error while updating account balance")

		return []transfer.Transfer{}, err
	}

	return transfers, nil
}
