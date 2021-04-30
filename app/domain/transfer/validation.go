package transfer

type Validation interface {
	Validate(Transfer) error
}

type InputCreationValidation interface {
	Validate(InputValue)
}
