package http

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/mellotonio/desafiogo/app/domain/account"
	usecasesAcc "github.com/mellotonio/desafiogo/app/domain/account/usecases"
	access "github.com/mellotonio/desafiogo/app/domain/authenticate"
	"github.com/mellotonio/desafiogo/app/domain/transfer"
	httpAccount "github.com/mellotonio/desafiogo/app/gateways/account"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/sirupsen/logrus"
)

// Presentation layer depends on Account, Transfer, Auth services
type API struct {
	AccountService  account.Service
	TransferService transfer.Service
	AuthService     access.Service
}

func newApi(AccountService account.Service, TransferService transfer.Service, AuthService access.Service) *API {
	return &API{
		AccountService:  AccountService,
		TransferService: TransferService,
		AuthService:     AuthService,
	}
}

func (a API) Start() {
	router := chi.NewRouter()

	accRepo := mem.NewAccountRepository(logrus.New())
	accServices := usecasesAcc.NewAccountService(accRepo)
	httpAccount.NewHandler(router, accServices)

	router.Use(httpAccount.NewHandler())



	fmt.Println("Starting api...")
	//err := http.ListenAndServe(":3000", v1)
	//log.Println(err)
}
