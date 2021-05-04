package account

import (
	"github.com/sirupsen/logrus"
)

type Validation interface {
	Validate() error
}

type InputCreationValidation interface {
	Validate(InputValue)
}

type ValidationStruct struct {
	account *Account
	log     *logrus.Entry
}
