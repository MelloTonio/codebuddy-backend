package card

import (
	"time"

	"github.com/mellotonio/desafiogo/app/infra/utils"
)

// Card Entity
type Card struct {
	Id         string    `json:"id"`
	DeckHolder string    `json:"deck_holder"`
	Owner      string    `json:"owner"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	Created_at time.Time `json:"created_at"`
}

func NewDeck(deckHolder string, owner string, question string, answer string) *Card {
	return &Card{
		Id:         utils.GenUUID(),
		DeckHolder: deckHolder,
		Owner:      owner,
		Question:   question,
		Answer:     answer,
		Created_at: time.Now(),
	}
}
