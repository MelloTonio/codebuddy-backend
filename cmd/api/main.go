package main

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	accountUsecases "github.com/mellotonio/desafiogo/app/domain/account/usecases"
	authUsecases "github.com/mellotonio/desafiogo/app/domain/authenticate/usecases"
	cardUsecases "github.com/mellotonio/desafiogo/app/domain/card/usecases"
	deckUsecases "github.com/mellotonio/desafiogo/app/domain/deck/usecases"

	TransferUsecases "github.com/mellotonio/desafiogo/app/domain/transfer/usecases"
	"github.com/mellotonio/desafiogo/app/gateways/http"
	mem "github.com/mellotonio/desafiogo/app/infra/persistence/memory"
	"github.com/mellotonio/desafiogo/app/infra/persistence/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load("../../.env")

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"postgres",
		"postgres",
		"desafiogo",
		"5432",
		"desafiogo")

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
	deckRepo := postgres.NewDeckRepository(db, log)
	cardRepo := postgres.NewCardRepository(db, log)

	// Services
	deckServices := deckUsecases.NewDeckUsecase(deckRepo, cardRepo)
	accountServices := accountUsecases.NewAccountService(accRepo)
	transferServices := TransferUsecases.NewTransfService(transfRepo, accRepo, trxRepo)
	authServices := authUsecases.NewAccessService(accRepo)
	cardServices := cardUsecases.NewCardUsecase(cardRepo)

	// API init
	API := http.NewApi(accountServices, transferServices, authServices, deckServices, cardServices)

	API.Start("0.0.0.0", "3001")
}

/*
docker-compose up --build



[{
    "owner":"567ce3f4-58e2-4bfd-b088-90f4ac2c056e",
    "question":"who is more fuck",
    "answer":"tefao"
},
{
    "owner":"567ce3f4-58e2-4bfd-b088-90f4ac2c056e",
    "question":"teste",
    "answer":"teste"
},
{
    "owner":"567ce3f4-58e2-4bfd-b088-90f4ac2c056e",
    "question":"teste",
    "answer":"teste"
}


]
*/
