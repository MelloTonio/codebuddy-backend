package deck

type Repository interface {
	Store(deckName string) (string, error)
}
