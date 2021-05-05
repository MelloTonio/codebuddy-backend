package mem

import (
	"database/sql/driver"

	"github.com/sirupsen/logrus"
)

func NewRepositoryTransaction() *MemRepositoryTrx {
	return &MemRepositoryTrx{log: logrus.New()}
}

type MemRepositoryTrx struct {
	log *logrus.Logger
}

func (MemRepositoryTrx) Commit() error {
	return nil
}

func (MemRepositoryTrx) Rollback() error {
	return nil
}

func (rr *MemRepositoryTrx) Tx() driver.Tx {
	return &MemRepositoryTrx{}
}
