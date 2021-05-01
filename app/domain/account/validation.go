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

/*func NewAccountValidation(acc *Account) Validation {
	return &ValidationStruct{
		account: acc,
		log:     logrus.NewEntry(logrus.New()).WithField("source", "NewAccountValidation"),
	}
}

func (vldStruct *ValidationStruct) Validate() error {
	if vldStruct.account.Id == "" {

		err := errors.EmptyAccountID_Err
		vldStruct.log.WithError(err).Error("Empty Account Id")
		return err
	}

	vldStruct.account.Cpf == "" {
		err := errors.EmptyAccountCPF_Err
		vldStruct.log.WithError(err).Error("Empty Account CPF")
		return err
	}

	return nil
}
*/
