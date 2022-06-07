package card

type Usecase interface {
	Store(*Card) error
}
