package card

type Usecase interface {
	Store(*Card) error
	Get(userID string) ([]Card, error)
	Delete(cardID string) error
	TradeCard(card, cardOwner, cardReceiver string) error
}
