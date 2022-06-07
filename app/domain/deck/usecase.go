package deck

import "github.com/mellotonio/desafiogo/app/domain/card"

type Usecase interface {
	CreateAndPopulateDeck([]card.Card) error
}

/*
When a significant process or transformation in the domain is not a natural responsibility
of an ENTITY or VALUE OBJECT, add an operation to the model as standalone interface declared as a SERVICE.
*/
