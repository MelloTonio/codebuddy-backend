package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	usecasesAcc "github.com/mellotonio/desafiogo/app/domain/account/usecases"
	usecasesAuth "github.com/mellotonio/desafiogo/app/domain/authenticate/usecases"
	usecasesTransf "github.com/mellotonio/desafiogo/app/domain/transfer/usecases"
	"github.com/mellotonio/desafiogo/app/gateways/http"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/sirupsen/logrus"
)

func main() {

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("dbname"))

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
	accRepo := mem.NewAccountRepository(log)
	transfRepo := mem.NewTransferRepository(log)
	trxRepo := mem.NewRepositoryTransaction()

	// Services
	accountServices := usecasesAcc.NewAccountService(accRepo)
	transferServices := usecasesTransf.NewTransfService(transfRepo, accRepo, trxRepo)
	authServices := usecasesAuth.NewAccessService(accRepo)

	// API init
	API := http.NewApi(accountServices, transferServices, authServices)

	API.Start("0.0.0.0", "3000")
}
