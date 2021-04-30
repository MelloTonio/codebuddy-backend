package mem

import (
	"testing"
	"time"

	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type testCaseSingle struct {
	Case          string
	In            account.Account
	ExpectedValue account.Account
	ExpectedError error
}

var (
	logger                  = logrus.New()
	NewMemAccountRepository = NewAccountRepository(logger)
)

var PassingStoreCases = []testCaseSingle{
	{
		Case: "Creating a valid account",
		In: account.Account{
			Id:         utils.GenUUID(),
			Name:       "Antonio Mello",
			Balance:    0,
			Cpf:        "12345678938",
			Secret:     "secretPass",
			Created_at: time.Now().Add(3 * time.Hour),
		},
		ExpectedError: nil,
	},
	{
		Case: "Creating a valid account",
		In: account.Account{
			Id:         utils.GenUUID(),
			Name:       "João Fernando",
			Balance:    0,
			Cpf:        "12345678939",
			Secret:     "secretPass_2",
			Created_at: time.Now().Add(10 * time.Hour),
		},
		ExpectedError: nil,
	},
}

var FailingStoreCases = []testCaseSingle{
	{
		Case: "Id cannot be empty",
		In: account.Account{
			Name:       "Antonio Mello",
			Balance:    0,
			Cpf:        "12345678938",
			Secret:     "secretPass",
			Created_at: time.Now().Add(3 * time.Hour),
		},
		ExpectedError: errors.EmptyAccountID_Err,
	},
	{
		Case: "Account Not Found",
		In: account.Account{
			Id:         utils.GenUUID(),
			Name:       "José Fora",
			Balance:    0,
			Cpf:        "xxx-xxx-xxx-xx",
			Secret:     "secretPass_2",
			Created_at: time.Now().Add(10 * time.Hour),
		},
		ExpectedError: errors.AccountNotFound_Err,
	},
}

func TestAccountRepo_Store_GetById_GetByCPF(t *testing.T) {

	t.Run("pass", func(t *testing.T) {

		t.Run("account.Repository.Store test", func(t *testing.T) {
			for i := range PassingStoreCases {
				t.Log(PassingStoreCases[i].Case)
				assert.NoError(t, NewMemAccountRepository.Store(&PassingStoreCases[i].In))
			}
		})

		t.Run("account.Repository.GetById test", func(t *testing.T) {
			for _, tc := range PassingStoreCases {
				t.Log(tc.Case)
				ac, err := NewMemAccountRepository.GetById(tc.In.Id)

				assert.NoError(t, err)

				assert.Equal(t, tc.In, ac)
			}
		})

		t.Run("account.Repository.GetByCPF test", func(t *testing.T) {
			for _, tc := range PassingStoreCases {
				t.Log(tc.Case)
				ac, err := NewMemAccountRepository.GetByCPF(tc.In.Cpf)

				assert.NoError(t, err)

				assert.Equal(t, tc.In, ac)
			}
		})

	})
}
