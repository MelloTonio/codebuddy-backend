package usecases

import "github.com/mellotonio/desafiogo/app/domain/card"

func (du CardUsecase) Get(userID string) ([]card.Card, error) {
	card, err := du.CardRepository.GetAllCards(userID)
	if err != nil {
		return nil, err
	}

	return card, nil
}
