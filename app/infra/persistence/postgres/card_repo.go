package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/mellotonio/desafiogo/app/domain/errors"
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
			  user_id,
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

func (repo cardRepository) GetAllCards(userID string) ([]card.Card, error) {
	cardArr := []card.Card{}

	stmt := `
		SELECT 
			id,
			question,
			answer
		from cards WHERE user_id=$1`

	rows, err := repo.DB.QueryContext(context.Background(), stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tempCard card.Card

		err = rows.Scan(
			&tempCard.Id,
			&tempCard.Question,
			&tempCard.Answer,
		)

		if err != nil {
			repo.log.WithError(err).Error("error while getting cards (by id)")
			if err == sql.ErrNoRows {
				err = errors.ErrAccountNotFound
			}
			return nil, err
		}

		cardArr = append(cardArr, tempCard)
	}

	if err != nil {
		repo.log.WithError(err).Error("Error to get card")
		return nil, err
	}

	return cardArr, nil
}

func (repo cardRepository) Delete(cardID string) error {
	stmt := `
		DELETE from cards where id = $1`

	_, err := repo.DB.Exec(stmt,
		cardID,
	)

	if err != nil {
		repo.log.WithError(err).Error("Error while deleting deck")
		return err
	}

	return nil
}

func (repo cardRepository) TradeCard(card, cardOwner, cardReceiver string) error {
	stmt := `UPDATE cards SET user_id = $1 where id = $2`

	_, err := repo.DB.Exec(stmt, cardReceiver, card)
	if err != nil {
		repo.log.WithError(err).Error("Error while trading card")
		return err
	}

	return nil
}
