package card

type Repository interface {
	Store(*Card) error
}
