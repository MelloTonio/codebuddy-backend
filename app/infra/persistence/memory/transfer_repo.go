package mem

import (
	"time"

	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/sirupsen/logrus"
)

type memTransferRepo struct {
	transfers []transfer.Transfer
	log       *logrus.Entry
}

func NewTransferRepository(logger *logrus.Logger) transfer.Repository {
	return &memTransferRepo{
		transfers: []transfer.Transfer{},
		log:       logger.WithField("source", "memAccountRepository"),
	}
}

func (memoryRepo *memTransferRepo) Store(transf *transfer.Transfer) error {
	if transf.Id == "" {
		err := errors.ErrEmptyTransferID
		memoryRepo.log.WithError(err).Error("Empty Transfer Id")
		return err
	}

	transf.Created_at = time.Now()

	memoryRepo.transfers = append(memoryRepo.transfers, *transf)

	return nil
}

func (memoryRepo *memTransferRepo) GetById(id string) (transfer.Transfer, error) {
	if id == "" {
		err := errors.ErrEmptyTransferID
		memoryRepo.log.WithError(err).Error("Empty Transfer Id")
		return transfer.Transfer{}, err
	}

	transferId := memoryRepo.find(id)

	if transferId == -1 {
		err := errors.ErrTransferNotFound
		memoryRepo.log.WithError(err).Error("Transfer not found")
		return transfer.Transfer{}, err
	}

	return memoryRepo.transfers[transferId], nil

}

func (memoryRepo *memTransferRepo) ListByAccId(id string) ([]transfer.Transfer, error) {
	if id == "" {
		err := errors.ErrEmptyAccountID
		memoryRepo.log.WithError(err).Error("Empty Account Id")
		return []transfer.Transfer{}, err
	}

	var transferCollection []transfer.Transfer

	for _, v := range memoryRepo.transfers {
		if v.Account_origin_id == id || v.Account_destination_id == id {
			transferCollection = append(transferCollection, v)
		}
	}

	if len(transferCollection) == 0 {
		memoryRepo.log.Error("No transfer found for this account")
	}

	return transferCollection, nil

}

func (memoryRepo *memTransferRepo) GenerateId() string {
	return utils.GenUUID()
}

func (memoryRepo *memTransferRepo) find(id string) int {
	for i, v := range memoryRepo.transfers {
		if v.Id == id {
			return i
		}
	}

	return -1
}
