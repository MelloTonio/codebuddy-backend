package deck

import (
	"time"

	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/mellotonio/desafiogo/app/infra/utils"
)

// Deck Entity
type Deck struct {
	Id         string      `json:"id"`
	DeckName   string      `json:"deck_name"`
	Cards      []card.Card `json:"cards"`
	Created_at time.Time   `json:"created_at"`
}

func NewDeck(deckName string) *Deck {
	return &Deck{
		Id:         utils.GenUUID(),
		DeckName:   deckName,
		Created_at: time.Now(),
	}
}
