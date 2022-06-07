package usecases

import (
	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/mellotonio/desafiogo/app/domain/deck"
)

type DeckUsecase struct {
	DeckRepository deck.Repository
	CardRepository card.Repository
}

func NewDeckUsecase(DeckRepository deck.Repository, CardRepository card.Repository) *DeckUsecase {
	return &DeckUsecase{
		DeckRepository: DeckRepository,
		CardRepository: CardRepository,
	}
}

func (du DeckUsecase) CreateAndPopulateDeck(cards []card.Card) error {
	deckID, err := du.DeckRepository.Store("deck_test")
	if err != nil {
		return err
	}

	for _, card := range cards {
		card.DeckHolder = deckID

		err = du.CardRepository.Store(&card)
		if err != nil {
			return err
		}

	}

	return nil
}
