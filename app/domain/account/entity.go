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
