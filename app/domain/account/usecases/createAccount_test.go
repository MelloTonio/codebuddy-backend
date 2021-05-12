package accountUsecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {

	t.Run("CreateAccount", func(t *testing.T) {

		err := accountServices.CreateAccount(&fake_account)
		assert.NoError(t, err)
	})

}
