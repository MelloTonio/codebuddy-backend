package TransferUsecases

import (
	"github.com/mellotonio/desafiogo/app/domain/account"
	accountUsecases "github.com/mellotonio/desafiogo/app/domain/account/usecases"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/sirupsen/logrus"
)

// To be used in tests
var (
	newLogrus              = logrus.New()
	memtransferRepo        = mem.NewTransferRepository(newLogrus)
	memaccountRepo         = mem.NewAccountRepository(newLogrus)
	memtransaction         = mem.NewRepositoryTransaction()
	NewAccountService      = accountUsecases.NewAccountService(memaccountRepo)
	NewtransferenceService = NewTransfService(memtransferRepo, memaccountRepo, memtransaction)
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
		log:         logrus.NewEntry(logrus.New()),
		accountRepo: AccountRepo,
		Trx:         mem.NewRepositoryTransaction(),
	}
}
