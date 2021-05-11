package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/sirupsen/logrus"
)

// ToDo: create a file for each func
// ToDo: test each func

type transferRepository struct {
	DB  *sql.DB
	log *logrus.Entry
	tx  driver.Tx
}

func NewTransferRepository(db *sql.DB, logger *logrus.Logger) transfer.Repository {
	return &transferRepository{
		DB:  db,
		log: logger.WithField("source", "PostgresTransferRepository"),
	}
}

func (repo transferRepository) Store(transf *transfer.Transfer) error {
	if transf.Id == "" {
		err := errors.ErrEmptyTransferID
		repo.log.WithError(err).Error("Empty Transfer Id")
		return err
	}

	stmt := `
		INSERT INTO transfers
			(	
				id,
				account_origin_id,
				account_destination_id,
				amount)
			VALUES ($1, $2, $3, $4)
		returning created_at`

	err := repo.DB.QueryRowContext(context.Background(), stmt,
		transf.Id,
		transf.Account_origin_id,
		transf.Account_destination_id,
		transf.Amount,
	).Scan(&transf.Created_at)

	if err != nil {
		err := errors.ErrCreatingAccount
		repo.log.WithError(err).Error("Error while creating transfer")
		return err
	}

	// ToDo return created_at
	return nil
}

func (repo transferRepository) ListByAccId(accID string) ([]transfer.Transfer, error) {
	const query = `
	SELECT 
		id,
		account_origin_id,
		account_destination_id,
		amount,
		created_at
	FROM 
		transfers
	WHERE account_origin_id=$1 OR account_destination_id=$1`

	transfers := []transfer.Transfer{}

	rows, err := repo.DB.QueryContext(context.Background(), query, accID)

	if err != nil {
		repo.log.WithError(err).Error("error while getting transfers (by id)")
		return []transfer.Transfer{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempTransf transfer.Transfer

		err = rows.Scan(
			&tempTransf.Id,
			&tempTransf.Account_origin_id,
			&tempTransf.Account_destination_id,
			&tempTransf.Amount,
			&tempTransf.Created_at,
		)

		if err != nil {
			repo.log.WithError(err).Error("error while getting transfers (by id)")
			if err == sql.ErrNoRows {
				err = errors.ErrAccountNotFound
			}
			return []transfer.Transfer{}, err
		}

		transfers = append(transfers, tempTransf)
	}

	return transfers, nil
}

func (repo transferRepository) GetById(transfID string) (transfer.Transfer, error) {
	if transfID == "" {
		err := errors.ErrEmptyTransferID
		repo.log.WithError(err).Error("Empty Account Id")
		return transfer.Transfer{}, err
	}

	query := `
		SELECT
			id,
			account_origin_id,
			account_destination_id,
			amount,
			created_at
		FROM
			transfers
		WHERE
			id=$1`

	var transf transfer.Transfer

	err := repo.DB.QueryRowContext(context.Background(), query, transfID).Scan(
		&transf.Id,
		&transf.Account_origin_id,
		&transf.Account_destination_id,
		&transf.Amount,
		&transf.Created_at,
	)

	if err != nil {
		err := errors.ErrAccountNotFound
		repo.log.WithError(err).Error("Account not found")
		return transfer.Transfer{}, err
	}

	return transf, nil
}

func (repo transferRepository) Transaction(tx driver.Tx) transfer.Repository {
	return &transferRepository{
		DB:  repo.DB,
		log: repo.log,
		tx:  tx,
	}
}

func (repo transferRepository) GenerateId() string {
	return utils.GenUUID()
}
