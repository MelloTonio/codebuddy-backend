package account

type Service interface {
	CreateAccount(*Account) error
	UpdateBalance(Account) error
	GetAccount(string) (Account, error)
	ShowAccounts() ([]Account, error)
}

/*
When a significant process or transformation in the domain is not a natural responsibility
of an ENTITY or VALUE OBJECT, add an operation to the model as standalone interface declared as a SERVICE.
*/
