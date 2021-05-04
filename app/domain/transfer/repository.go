package transfer

import "database/sql/driver"

type Repository interface {
	Store(*Transfer) error
	GetById(string) (Transfer, error)
	ListByAccId(string) ([]Transfer, error)
	GenerateId() string
	Transaction(driver.Tx) Repository
}
