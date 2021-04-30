package account

type Validation interface {
	Validate(Account) error
}

type InputCreationValidation interface {
	Validate(InputValue)
}
