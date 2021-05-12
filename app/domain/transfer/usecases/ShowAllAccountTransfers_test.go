package TransferUsecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ShowAllAccountTransfersTest(t *testing.T) {
	for _, v := range PassingStoreCases {
		NewAccountService.CreateAccount(&v)
	}

	_, err := NewtransferenceService.Transfer(PassingStoreCases[0].Id, PassingStoreCases[1].Id, 2000)
	assert.NoError(t, err)

	_, err = NewAccountService.ShowAccounts()

	assert.NoError(t, err)

}
