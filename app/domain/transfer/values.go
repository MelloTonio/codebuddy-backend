package transfer

type InputValue struct {
	Account_origin_id      string `json:"account_origin_id"`
	Account_destination_id string `json:"account_destination_id"`
	Amount                 int    `json:"amount"`
}

// Validates the input value
func (iptVl InputValue) Validate(validationStrategy InputCreationValidation) error {
	return iptVl.Validate(validationStrategy)
}

type BalanceValue struct {
	Balance int `json:"balance"`
}
