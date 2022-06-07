package usecases

import (
	"github.com/mellotonio/desafiogo/app/domain/card"
)

type CardUsecase struct {
	CardRepository card.Repository
}

func NewCardUsecase(CardRepository card.Repository) *CardUsecase {
	return &CardUsecase{
		CardRepository: CardRepository,
	}
}

func (du CardUsecase) Store(card *card.Card) error {
	err := du.CardRepository.Store(card)
	if err != nil {
		return err
	}

	return nil
}
