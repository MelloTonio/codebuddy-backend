package account

type InputValue struct {
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

// Validates the input value
func (iptVl InputValue) Validate(validationStrategy InputCreationValidation) error {
	return iptVl.Validate(validationStrategy)
}

type BalanceValue struct {
	Balance int `json:"balance"`
}
