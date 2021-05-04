package transferUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/sirupsen/logrus"
)

type Services struct {
	log         *logrus.Entry
	transfRepo  transfer.Repository
	accountRepo account.Repository
	Trx         *mem.MemRepositoryTrx
}

func NewTransfService(
	TransferRepo transfer.Repository,
	AccountRepo account.Repository,
	Trx *mem.MemRepositoryTrx,
) transfer.Service {

	return &Services{
		transfRepo:  TransferRepo,
		log:         logrus.NewEntry(&logrus.Logger{}),
		accountRepo: AccountRepo,
		Trx:         mem.NewRepositoryTransaction(),
	}
}
