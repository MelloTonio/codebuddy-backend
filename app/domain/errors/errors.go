package errors

import (
	"errors"
	"fmt"
)

// Domain errors.
var (
	ErrDomain               = errors.New("domain Error")
	ErrEmptyAccountID       = fmt.Errorf("%w: %v", ErrDomain, "Id cannot be Empty")
	ErrEmptyTransferID      = fmt.Errorf("%w: %v", ErrDomain, "Id cannot be Empty")
	ErrAccountNotFound      = fmt.Errorf("%w: %v", ErrDomain, "Account Not Found")
	ErrEmptyCPF             = fmt.Errorf("%w: %v", ErrDomain, "CPF cannot be empty")
	ErrAccountAlreadyExists = fmt.Errorf("%w: %v", ErrDomain, "Account already exists")
	ErrTransferNotFound     = fmt.Errorf("%w: %v", ErrDomain, "Transfer not found")
	ErrInsufficienteBalance = fmt.Errorf("%w: %v", ErrDomain, "Insuficient balance")
	ErrPasswordsDontMatch   = fmt.Errorf("%w: %v", ErrDomain, "Passwords don't match")
	ErrEmptyAccountSecret   = fmt.Errorf("%w: %v", ErrDomain, "Empty Account Secret")
	ErrSigningJwt           = fmt.Errorf("%w: %v", ErrDomain, "Error while signing JWT")
	ErrCreatingAccount      = fmt.Errorf("%w: %v", ErrDomain, "Error while creating account")
	ErrUpdatingBalance      = fmt.Errorf("%w: %v", ErrDomain, "Error while updating balance")
	ErrUnauthorized         = fmt.Errorf("%w: %v", ErrDomain, "Unauthorized")
)
