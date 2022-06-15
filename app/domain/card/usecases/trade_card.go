package usecases

import "fmt"

func (du CardUsecase) TradeCard(card, cardOwner, cardReceiver string) error {
	fmt.Println(card, cardOwner, cardReceiver)
	err := du.CardRepository.TradeCard(card, cardOwner, cardReceiver)
	if err != nil {
		return err
	}

	return nil
}
