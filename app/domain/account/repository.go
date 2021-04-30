package account

// Manage the persistence related to account entity
type Repository interface {
	Store(*Account) error
	GetBalance(Account) (int, error)
	ExistsByCPF(*Account) (bool, error)
	UpdateBalance(Account, int) error
	GetById(string) (Account, error)
	GetByCPF(string) (Account, error)
	ShowAll() ([]Account, error)
	GenerateID() string
	// Transaction(driver.Tx) Repository
}
