package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	usecasesAcc "github.com/mellotonio/desafiogo/app/domain/account/usecases"
	usecasesAuth "github.com/mellotonio/desafiogo/app/domain/authenticate/usecases"
	usecasesTransf "github.com/mellotonio/desafiogo/app/domain/transfer/usecases"
	"github.com/mellotonio/desafiogo/app/gateways/http"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/mellotonio/desafiogo/app/infra/persistence/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"postgres", "postgres", "localhost", "5432", "desafiogo")

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	log := logrus.New()

	// Repositories
	accRepo := postgres.NewAccountRepository(db, log)
	transfRepo := postgres.NewTransferRepository(db, log)
	trxRepo := mem.NewRepositoryTransaction() // ToDo: postgres trx repo

	// Services
	accountServices := usecasesAcc.NewAccountService(accRepo)
	transferServices := usecasesTransf.NewTransfService(transfRepo, accRepo, trxRepo)
	authServices := usecasesAuth.NewAccessService(accRepo)

	// API init
	API := http.NewApi(accountServices, transferServices, authServices)

	API.Start("0.0.0.0", "3001")
}
