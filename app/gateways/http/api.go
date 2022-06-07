package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/account"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/mellotonio/desafiogo/app/domain/deck"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	httpCards "github.com/mellotonio/desafiogo/app/gateways/card"
	httpDeck "github.com/mellotonio/desafiogo/app/gateways/decks"
)

// Presentation layer depends on Account, Transfer, Auth services
type API struct {
	AccountService  account.Service
	TransferService transfer.Service
	AuthService     access.Service
	DeckService     deck.Usecase
	CardService     card.Usecase
}

func NewApi(AccountService account.Service, TransferService transfer.Service, AuthService access.Service, DeckService deck.Usecase, CardService card.Usecase) *API {
	return &API{
		AccountService:  AccountService,
		TransferService: TransferService,
		AuthService:     AuthService,
		DeckService:     DeckService,
		CardService:     CardService,
	}
}

func (api API) Start(host string, port string) {
	router := chi.NewMux()

	// Handlers - Account & Transfer
	httpDeck.NewHandler(router, api.DeckService)
	httpCards.NewHandler(router, api.CardService)

	applicationPort := fmt.Sprintf("%s:%s", host, port)

	server := &http.Server{
		Handler:      router,
		Addr:         applicationPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	fmt.Println("Starting api...")
	log.Fatal(server.ListenAndServe())
}
