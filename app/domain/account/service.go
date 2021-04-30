package account

/*
When a significant process or transformation in the domain is not a natural responsibility
of an ENTITY or VALUE OBJECT, add an operation to the model as standalone interface declared as a SERVICE.
*/

type Service interface {
	CreateAccount(*Account) error
	GetBalance(Account) (BalanceValue, error)
	UpdateBalance(Account) error
	GetAccount(string) (Account, error)
	ShowAccounts() ([]Account, error)
}
