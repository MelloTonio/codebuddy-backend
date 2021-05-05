package usecasesAcc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Depends on Create Account test (run package tests)
func TestShowAccounts(t *testing.T) {

	t.Run("ShowAccounts", func(t *testing.T) {

		accounts, err := accountServices.ShowAccounts()

		for _, v := range accounts {
			t.Logf("Name: %s - Balance: %d", v.Name, v.Balance)
		}

		assert.NoError(t, err)
	})
}
