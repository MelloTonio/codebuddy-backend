package TransferUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
)

func (transfSvc *Services) Transfer(Account_origin_id string, Account_destination_id string, amount int) (transfer.Transfer, error) {

	if amount < 1 {
		err := errors.ErrNegativeAmount
		transfSvc.log.WithError(err).Error("Amount cannot < 1")
		return transfer.Transfer{}, err
	}

	accountRepo := transfSvc.accountRepo

	origin_Account, err := accountRepo.GetById(Account_origin_id)

	if err != nil {
		transfSvc.log.WithError(err).Error("Account 'Origin' not found")
		return transfer.Transfer{}, err
	}

	destination_Account, err := accountRepo.GetById(Account_destination_id)

	if err != nil {
		transfSvc.log.WithError(err).Error("Account 'Destination' not found")
		return transfer.Transfer{}, err
	}

	if origin_Account.Balance < amount {
		err := errors.ErrInsufficienteBalance

		transfSvc.log.WithError(err).Error("Not enough balance Account 'Destination'")
		return transfer.Transfer{}, err
	}

	origin_Account.Balance -= amount
	destination_Account.Balance += amount

	tx := transfSvc.Trx

	err = accountRepo.Transaction(tx).UpdateBalance(&origin_Account)

	if err != nil {
		transfSvc.log.WithError(err).Error("failed to update the balance of the origin Account")
		tx.Rollback()
		return transfer.Transfer{}, err
	}

	err = accountRepo.Transaction(tx).UpdateBalance(&destination_Account)

	if err != nil {
		transfSvc.log.WithError(err).Error("failed to update the balance of the destination Account")
		tx.Rollback()
		return transfer.Transfer{}, err
	}

	transfRepo := transfSvc.transfRepo

	newTransfer := transfer.Transfer{
		Id:                     transfRepo.GenerateId(),
		Account_origin_id:      origin_Account.Id,
		Account_destination_id: destination_Account.Id,
		Amount:                 amount,
	}

	err = transfRepo.Transaction(tx).Store(&newTransfer)

	if err != nil {
		transfSvc.log.WithError(err).Error("Failed to store new transaction")
		tx.Rollback()
		return transfer.Transfer{}, err
	}

	err = tx.Commit()

	if err != nil {
		transfSvc.log.WithError(err).Error("Failed to finish the transaction")
		return transfer.Transfer{}, err
	}

	return newTransfer, nil
}
