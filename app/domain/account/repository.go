package account

import "database/sql/driver"

// Manage the persistence related to account entity
type Repository interface {
	Store(*Account) error
	GetBalance(Account) error
	ExistsByCPF(Account) error
	UpdateBalance(Account) error
	GetByCPF(string) (Account, error)
	ShowAll() ([]Account, error)
	GeneratedID() string
	Transaction(driver.Tx) Repository
}
