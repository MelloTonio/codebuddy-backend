package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/account"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	httpAccount "github.com/mellotonio/desafiogo/app/gateways/account"
	httpAuth "github.com/mellotonio/desafiogo/app/gateways/auth"
	httpTransfer "github.com/mellotonio/desafiogo/app/gateways/transfer"
)

// Presentation layer depends on Account, Transfer, Auth services
type API struct {
	AccountService  account.Service
	TransferService transfer.Service
	AuthService     access.Service
}

func NewApi(AccountService account.Service, TransferService transfer.Service, AuthService access.Service) *API {
	return &API{
		AccountService:  AccountService,
		TransferService: TransferService,
		AuthService:     AuthService,
	}
}

func (api API) Start(host string, port string) {
	router := chi.NewMux()

	// Handlers - Account & Transfer
	httpAccount.NewHandler(router, api.AccountService)
	httpTransfer.NewHandler(router, api.TransferService)
	httpAuth.NewHandler(router, api.AuthService)

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
