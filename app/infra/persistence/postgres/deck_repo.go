package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/mellotonio/desafiogo/app/domain/deck"
	"github.com/sirupsen/logrus"
)

type deckRepository struct {
	DB  *sql.DB
	log *logrus.Entry
	tx  driver.Tx
}

func NewDeckRepository(db *sql.DB, logger *logrus.Logger) deck.Repository {
	return &deckRepository{
		DB:  db,
		log: logger.WithField("source", "PostgresDeckRepository"),
	}
}

func (repo deckRepository) Store(deckName string) (string, error) {
	var deckID string

	stmt := `
		INSERT INTO decks
			(	
			 deck_name
			)
			VALUES ($1)
		returning id`

	err := repo.DB.QueryRowContext(context.Background(), stmt,
		deckName,
	).Scan(&deckID)

	if err != nil {
		repo.log.WithError(err).Error("Error while creating deck")
		return "", err
	}

	return deckID, nil
}
