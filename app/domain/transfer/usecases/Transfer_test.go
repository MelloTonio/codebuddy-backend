package TransferUsecases

import (
	"testing"
	"time"

	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/stretchr/testify/assert"
)

var PassingStoreCases = []account.Account{
	{
		Id:         utils.GenUUID(),
		Name:       "Antonio Mello",
		Balance:    5000,
		Cpf:        "12345678938",
		Secret:     "secretPass",
		Created_at: time.Now().Add(3 * time.Hour),
	},
	{Id: utils.GenUUID(),
		Name:       "João Fernando",
		Balance:    3000,
		Cpf:        "12345678939",
		Secret:     "secretPass_2",
		Created_at: time.Now().Add(10 * time.Hour),
	},
}

func TestTransfer(t *testing.T) {
	t.Run("Transfer", func(t *testing.T) {

		for _, v := range PassingStoreCases {
			NewAccountService.CreateAccount(&v)
		}

		_, err := NewtransferenceService.Transfer(PassingStoreCases[0].Id, PassingStoreCases[1].Id, 2000)
		assert.NoError(t, err)

		accounts, err := NewAccountService.ShowAccounts()

		assert.NoError(t, err)

		assert.NotEqual(t, accounts[0].Balance, 5000)
		assert.NotEqual(t, accounts[1].Balance, 3000)

		assert.Equal(t, accounts[0].Balance, 3000)
		assert.Equal(t, accounts[1].Balance, 5000)
	})

	t.Run("Transfer with negative values", func(t *testing.T) {

		for _, v := range PassingStoreCases {
			NewAccountService.CreateAccount(&v)
		}

		_, err := NewtransferenceService.Transfer(PassingStoreCases[0].Id, PassingStoreCases[1].Id, -2000)
		assert.Error(t, err)

	})
}
