package errors

import (
	"errors"
	"fmt"
)

// Domain errors.
var (
	Domain_Err          = errors.New("Domain Error")
	EmptyAccountID_Err  = fmt.Errorf("%w: %v", Domain_Err, "Id cannot be Empty")
	AccountNotFound_Err = fmt.Errorf("%w: %v", Domain_Err, "Account Not Found")
	EmptyCPF_Err        = fmt.Errorf("%w: %v", Domain_Err, "CPF cannot be empty")
)
