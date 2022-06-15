package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/mellotonio/desafiogo/app/domain/errors"
	"github.com/mellotonio/desafiogo/app/domain/user"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	DB  *sql.DB
	log *logrus.Entry
	tx  driver.Tx
}

func NewUserRepository(db *sql.DB, logger *logrus.Logger) user.Repository {
	return &userRepository{
		DB:  db,
		log: logger.WithField("source", "PostgresDeckRepository"),
	}
}

func (repo userRepository) GetAll() ([]user.User, error) {
	userArr := []user.User{}

	stmt := `
		SELECT 
			id,
			nickname,
			email
		from users`

	rows, err := repo.DB.QueryContext(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tempUser user.User

		err = rows.Scan(
			&tempUser.Id,
			&tempUser.Nickname,
			&tempUser.Email,
		)

		if err != nil {
			repo.log.WithError(err).Error("error while getting cards (by id)")
			if err == sql.ErrNoRows {
				err = errors.ErrAccountNotFound
			}
			return nil, err
		}

		userArr = append(userArr, tempUser)
	}

	if err != nil {
		repo.log.WithError(err).Error("Error to get card")
		return nil, err
	}

	return userArr, nil
}

func (repo userRepository) Get(nickname string) (user.User, error) {
	tempUser := user.User{}

	stmt := `
		SELECT 
			id
		from users where nickname = $1`

	err := repo.DB.QueryRowContext(context.Background(), stmt, nickname).Scan(
		&tempUser.Id,
	)
	if err != nil {
		return user.User{}, err
	}

	return tempUser, nil
}
