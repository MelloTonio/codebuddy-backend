package card

type Repository interface {
	Store(*Card) error
	GetAllCards(userID string) ([]Card, error)
	Delete(cardID string) error
	TradeCard(card, cardOwner, cardReceiver string) error
}
