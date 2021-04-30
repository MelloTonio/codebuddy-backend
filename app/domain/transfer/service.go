package transfer

type Service interface {
	Transfer(Account_origin_id string, Account_destination_id string, amount int) (Transfer, error)
	ShowAllAccountTransfers(Account_Id string) ([]Transfer, error)
}
