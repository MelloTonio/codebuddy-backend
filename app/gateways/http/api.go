package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/mellotonio/desafiogo/app/domain/account"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/domain/card"
	"github.com/mellotonio/desafiogo/app/domain/deck"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	"github.com/mellotonio/desafiogo/app/domain/user"
	httpCards "github.com/mellotonio/desafiogo/app/gateways/card"
	httpDeck "github.com/mellotonio/desafiogo/app/gateways/decks"
	httpUser "github.com/mellotonio/desafiogo/app/gateways/user"
)

// Presentation layer depends on Account, Transfer, Auth services
type API struct {
	AccountService  account.Service
	TransferService transfer.Service
	AuthService     access.Service
	DeckService     deck.Usecase
	UserService     user.Usecase
	CardService     card.Usecase
}

func NewApi(AccountService account.Service, TransferService transfer.Service, AuthService access.Service, DeckService deck.Usecase, CardService card.Usecase, UserService user.Usecase) *API {
	return &API{
		AccountService:  AccountService,
		TransferService: TransferService,
		AuthService:     AuthService,
		DeckService:     DeckService,
		UserService:     UserService,
		CardService:     CardService,
	}
}

func (api API) Start(host string, port string) {
	router := chi.NewMux()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Handlers - Account & Transfer
	httpDeck.NewHandler(router, api.DeckService)
	httpCards.NewHandler(router, api.CardService)
	httpUser.NewHandler(router, api.UserService)

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
