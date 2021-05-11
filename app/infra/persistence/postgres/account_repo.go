package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/mellotonio/desafiogo/app/domain/account"
	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/sirupsen/logrus"
)

// ToDo: create a file for each func
// ToDo: test each func

type accountRepository struct {
	DB  *sql.DB
	log *logrus.Entry
	tx  driver.Tx
}

func NewAccountRepository(db *sql.DB, logger *logrus.Logger) account.Repository {
	return &accountRepository{
		DB:  db,
		log: logger.WithField("source", "PostgresAccountRepository"),
	}
}

func (repo accountRepository) Store(account *account.Account) error {
	if account.Id == "" {
		err := errors.ErrEmptyAccountID
		repo.log.WithError(err).Error("Empty Account Id")
		return err
	}

	stmt := `
		INSERT INTO accounts
			(	id,
				name,
				cpf,
				secret,
				balance)
			VALUES ($1, $2, $3, $4, $5)
		returning created_at`

	err := repo.DB.QueryRowContext(context.Background(), stmt,
		account.Id,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance,
	).Scan(&account.Created_at)

	if err != nil {
		err := errors.ErrCreatingAccount
		repo.log.WithError(err).Error("Error while creating account")
		return err
	}

	// ToDo return created_at
	return nil
}

func (repo accountRepository) UpdateBalance(acc *account.Account) error {
	const stmt = `
	UPDATE accounts
		SET balance = $1
	WHERE
		id = $2`

	_, err := repo.DB.Exec(stmt, acc.Balance, acc.Id)

	if err != nil {
		err := errors.ErrUpdatingBalance
		repo.log.WithError(err).Error("Error while updating balance")
		return err
	}

	return nil
}

// ToDO remove getBalance from all repositories/signatures
func (repo accountRepository) GetBalance(acc account.Account) (int, error) {
	if acc.Id == "" {
		err := errors.ErrEmptyAccountID
		repo.log.WithError(err).Error("Empty Account Id")
		return -1, err
	}

	query := `
	SELECT
		balance
	FROM
		accounts
	WHERE
		id=$1`

	var tempAcc account.Account

	err := repo.DB.QueryRowContext(context.Background(), query, acc.Id).Scan(
		&tempAcc.Balance,
	)

	if err != nil {
		err := errors.ErrAccountNotFound
		repo.log.WithError(err).Error("Account Not Found")
		return -1, err
	}

	return tempAcc.Balance, nil
}

func (repo accountRepository) ShowAll() ([]account.Account, error) {

	const query = `
	SELECT
		id,
		name,
		cpf,
		balance,
		created_at
	FROM
		accounts
	ORDER BY
		created_at
	`
	accounts := []account.Account{}

	rows, err := repo.DB.QueryContext(context.Background(), query)

	for rows.Next() {
		var tempAcc account.Account

		err = rows.Scan(
			&tempAcc.Id,
			&tempAcc.Name,
			&tempAcc.Cpf,
			&tempAcc.Balance,
			&tempAcc.Created_at,
		)

		if err != nil {
			repo.log.WithError(err).Error("error while getting all accounts")
			if err == sql.ErrNoRows {
				err = errors.ErrAccountNotFound
			}
			return []account.Account{}, err
		}

		accounts = append(accounts, tempAcc)
	}

	return accounts, nil
}

func (repo accountRepository) GetById(accId string) (account.Account, error) {
	if accId == "" {
		err := errors.ErrEmptyAccountID
		repo.log.WithError(err).Error("Empty Account Id")
		return account.Account{}, err
	}

	query := `
		SELECT
			id,
			name,
			cpf,
			balance,
			created_at
		FROM
			accounts
		WHERE
			id=$1`

	var acc account.Account

	err := repo.DB.QueryRowContext(context.Background(), query, accId).Scan(
		&acc.Id,
		&acc.Name,
		&acc.Cpf,
		&acc.Balance,
		&acc.Created_at,
	)

	if err != nil {
		err := errors.ErrAccountNotFound
		repo.log.WithError(err).Error("Account not found")
		return account.Account{}, err
	}

	return acc, nil
}

func (repo accountRepository) GetByCPF(accCPF string) (account.Account, error) {
	if accCPF == "" {
		err := errors.ErrEmptyCPF
		repo.log.WithError(err).Error("Empty Account CPF")
		return account.Account{}, err
	}

	query := `
		SELECT
			id,
			name,
			cpf,
			secret,
			balance,
			created_at
		FROM
			accounts
		WHERE
			cpf=$1`

	var acc account.Account

	err := repo.DB.QueryRowContext(context.Background(), query, accCPF).Scan(
		&acc.Id,
		&acc.Name,
		&acc.Cpf,
		&acc.Secret,
		&acc.Balance,
		&acc.Created_at,
	)

	if err != nil {
		err := errors.ErrAccountNotFound
		repo.log.WithError(err).Error("Account not found")
		return account.Account{}, err
	}

	return acc, nil
}

func (repo accountRepository) ExistsByCPF(acc *account.Account) (bool, error) {
	query := `
		SELECT
			id,
			name,
			cpf,
			balance,
			created_at
		FROM
			accounts
		WHERE
			cpf=$1`

	var newAcc account.Account

	err := repo.DB.QueryRowContext(context.Background(), query, acc.Cpf).Scan(
		&newAcc.Id,
		&newAcc.Name,
		&newAcc.Cpf,
		&newAcc.Balance,
		&newAcc.Created_at,
	)

	// Err == Not found
	if err != nil {
		err := errors.ErrAccountAlreadyExists
		repo.log.WithError(err).Error("Account already exists")
		return false, err
	}

	return true, nil
}

func (repo accountRepository) Transaction(tx driver.Tx) account.Repository {
	return &accountRepository{
		DB:  repo.DB,
		log: repo.log,
		tx:  tx,
	}
}

func (repo accountRepository) GenerateID() string {
	return utils.GenUUID()
}
