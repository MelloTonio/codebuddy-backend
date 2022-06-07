package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/sirupsen/logrus"
)

type cardRepository struct {
	DB  *sql.DB
	log *logrus.Entry
	tx  driver.Tx
}

func NewCardRepository(db *sql.DB, logger *logrus.Logger) card.Repository {
	return &cardRepository{
		DB:  db,
		log: logger.WithField("source", "PostgresCardRepository"),
	}
}

func (repo cardRepository) Store(Card *card.Card) error {
	var deckID string

	stmt := `
		INSERT INTO cards
			(	
			  deck_holder,
			  owner_id,
			  question,
		      answer
			)
			VALUES ($1, $2, $3, $4)
		returning id`

	fmt.Println(Card)
	err := repo.DB.QueryRowContext(context.Background(), stmt,
		Card.DeckHolder, Card.Owner, Card.Question, Card.Answer,
	).Scan(&deckID)

	if err != nil {
		repo.log.WithError(err).Error("Error while creating card")
		return err
	}

	return nil
}
