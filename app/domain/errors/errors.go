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
)
