package usecases

func (du CardUsecase) Delete(cardID string) error {
	err := du.CardRepository.Delete(cardID)
	if err != nil {
		return err
	}

	return nil
}
