package transfer

import "time"

type Transfer struct {
	Id                     string    `json:"id"`
	Account_origin_id      string    `json:"account_origin_id"`
	Account_destination_id string    `json:"account_destination_id"`
	Amount                 int       `json:"amount"`
	Created_at             time.Time `json:"created_at"`
}

func NewTransfer(Account_origin_id string, Account_destination_id string, Amount int) *Transfer {
	return &Transfer{
		Account_origin_id:      Account_origin_id,
		Account_destination_id: Account_destination_id,
		Amount:                 Amount,
	}
}
