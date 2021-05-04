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
			Cpf:        "12345678931",
			Secret:     "secretPass",
			Created_at: time.Now().Add(3 * time.Hour),
		},
		ExpectedError: errors.ErrEmptyAccountID,
	},
	{
		Case: "Account Not Found",
		In: account.Account{
			Id:         utils.GenUUID(),
			Name:       "José Fora",
			Balance:    0,
			Cpf:        "12345678932",
			Secret:     "secretPass_2",
			Created_at: time.Now().Add(10 * time.Hour),
		},
		ExpectedError: errors.ErrAccountNotFound,
	},
	{
		Case: "Empty CPF value",
		In: account.Account{
			Id:         utils.GenUUID(),
			Name:       "José Fora",
			Balance:    0,
			Secret:     "12345678933",
			Created_at: time.Now().Add(10 * time.Hour),
		},
		ExpectedError: errors.ErrEmptyCPF,
	},
	{
		Case: "Account already exists",
		In: account.Account{
			Id:         utils.GenUUID(),
			Name:       "José Fora",
			Cpf:        "12345678934",
			Balance:    0,
			Secret:     "secretPass_2",
			Created_at: time.Now().Add(10 * time.Hour),
		},
		ExpectedError: errors.ErrAccountAlreadyExists,
	},
}

func TestAccountRepo(t *testing.T) {

	t.Run("pass", func(t *testing.T) {

		t.Run("account.Repository.Store test", func(t *testing.T) {
			for i := range PassingStoreCases {
				t.Log(PassingStoreCases[i].Case)
				assert.NoError(t, NewMemAccountRepository.Store(&PassingStoreCases[i].In))
			}
		})

		t.Run("account.Repository.ShowAll test", func(t *testing.T) {
			accounts, err := NewMemAccountRepository.ShowAll()

			assert.NoError(t, err)

			assert.Equal(t, len(accounts), len(PassingStoreCases))
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

		t.Run("account.Repository.UpdateBalance test", func(t *testing.T) {
			for i, tc := range PassingStoreCases {

				t.Log(tc.Case)
				err := NewMemAccountRepository.UpdateBalance(&PassingStoreCases[i].In)
				assert.NoError(t, err)

				accounts, err := NewMemAccountRepository.ShowAll()
				assert.NoError(t, err)

				assert.Equal(t, accounts[i].Balance, PassingStoreCases[i].In.Balance)
			}
		})

		t.Run("account.Repository.getBalance test", func(t *testing.T) {
			for i, tc := range PassingStoreCases {

				t.Log(tc.Case)

				balance, err := NewMemAccountRepository.GetBalance(tc.In)
				assert.NoError(t, err)

				accounts, err := NewMemAccountRepository.ShowAll()
				assert.NoError(t, err)

				assert.Equal(t, accounts[i].Balance, balance)
			}
		})

		t.Run("account.Repository.existsByCPF test", func(t *testing.T) {
			for i, tc := range PassingStoreCases {

				t.Log(tc.Case)

				exists, err := NewMemAccountRepository.ExistsByCPF(&PassingStoreCases[i].In)
				assert.NoError(t, err)

				assert.True(t, exists)
			}
		})

	})

	t.Run("fail", func(t *testing.T) {

		t.Run("account.Repository.Store test", func(t *testing.T) {

			err := NewMemAccountRepository.Store(&FailingStoreCases[0].In)
			assert.Error(t, err)

			// Creating duplicate accounts to simulate AccountAlreadyExists error
			NewMemAccountRepository.Store(&FailingStoreCases[3].In)
			err = NewMemAccountRepository.Store(&FailingStoreCases[3].In)

			if assert.Error(t, err) {
				assert.Equal(t, errors.ErrAccountAlreadyExists, err)
			}

		})

		t.Run("account.Repository.GetById test", func(t *testing.T) {

			t.Log(FailingStoreCases[1].Case)

			// Search for a account that doesn't exist
			_, err := NewMemAccountRepository.GetById(FailingStoreCases[1].In.Id)

			if assert.Error(t, err) {
				assert.Equal(t, errors.ErrAccountNotFound, err)
			}

			// Search for a account with a invalid Id
			_, err = NewMemAccountRepository.GetById("")

			if assert.Error(t, err) {
				assert.Equal(t, errors.ErrEmptyAccountID, err)
			}

		})

		t.Run("account.Repository.GetByCPF test", func(t *testing.T) {

			// Search for a account without a CPF
			_, err := NewMemAccountRepository.GetByCPF(FailingStoreCases[2].In.Cpf)

			if assert.Error(t, err) {
				assert.Equal(t, errors.ErrEmptyCPF, err)
			}

			// Search for a account that doesn't exist
			_, err = NewMemAccountRepository.GetByCPF(FailingStoreCases[1].In.Cpf)

			if assert.Error(t, err) {
				assert.Equal(t, errors.ErrAccountNotFound, err)
			}

		})

		t.Run("account.Repository.ExistsByCPF test", func(t *testing.T) {

			// Search for a account without a ID
			_, err := NewMemAccountRepository.ExistsByCPF(&FailingStoreCases[0].In)

			if assert.Error(t, err) {
				assert.Equal(t, errors.ErrEmptyAccountID, err)
			}

		})

	})
}
