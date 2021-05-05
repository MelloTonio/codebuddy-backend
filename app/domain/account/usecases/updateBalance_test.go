package usecasesAcc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Depends on Create Account test (run package tests)
func TestUpdateBalance(t *testing.T) {
	t.Run("UpdateBalance", func(t *testing.T) {
		oldBalance := fake_account.Balance
		newBalance := 99999

		fake_account.Balance = newBalance

		err := accountServices.UpdateBalance(fake_account)

		assert.NoError(t, err)

		assert.NotEqual(t, fake_account.Balance, oldBalance)
		assert.Equal(t, newBalance, fake_account.Balance)

	})
}
