package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	accountUsecases "github.com/mellotonio/desafiogo/app/domain/account/usecases"
	authUsecases "github.com/mellotonio/desafiogo/app/domain/authenticate/usecases"
	TransferUsecases "github.com/mellotonio/desafiogo/app/domain/transfer/usecases"
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
		os.Getenv("DATABASE_NAME"),
		os.Getenv("PORT"),
		os.Getenv("DATABASE_NAME"))

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
	accountServices := accountUsecases.NewAccountService(accRepo)
	transferServices := TransferUsecases.NewTransfService(transfRepo, accRepo, trxRepo)
	authServices := authUsecases.NewAccessService(accRepo)

	// API init
	API := http.NewApi(accountServices, transferServices, authServices)

	API.Start("0.0.0.0", "3001")
}
