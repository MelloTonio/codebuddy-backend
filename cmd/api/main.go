package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	godotenv.Load("../../.env.example")

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DB_NAME"))

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
	trxRepo := mem.NewRepositoryTransaction()

	// Services
	accountServices := usecasesAcc.NewAccountService(accRepo)
	transferServices := usecasesTransf.NewTransfService(transfRepo, accRepo, trxRepo)
	authServices := usecasesAuth.NewAccessService(accRepo)

	// API init
	API := http.NewApi(accountServices, transferServices, authServices)

	API.Start("0.0.0.0", "3001")
}
