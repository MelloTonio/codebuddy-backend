package account

import "time"

// Account Entity
type Account struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Cpf        string    `json:"cpf"`
	Secret     string    `json:"secret"`
	Balance    int       `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

func NewAccount(name string, cpf string, secret string, balance int) *Account {
	return &Account{
		// ID:     ,		// ToDo: Function to generate ID
		Name:       name,
		Cpf:        cpf, // ToDo: Function to parse CPF
		Secret:     secret,
		Balance:    balance,
		Created_at: time.Now(),
	}
}
