package accountUsecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Depends on Create Account test (run package tests)
func TestGetAccount(t *testing.T) {
	t.Run("GetAccount", func(t *testing.T) {

		_, err := accountServices.GetAccount("555-444-333")
		assert.NoError(t, err)
	})

}
