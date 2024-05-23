package main

import (
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	localPool "github.com/mellotonio/desafiogo"

	"github.com/mellotonio/desafiogo/app/domain/studyGroups/services"
	"github.com/mellotonio/desafiogo/app/gateways/http"
	mongodb "github.com/mellotonio/desafiogo/app/infra/persistence/mongoDB"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load("../../.env")

	log := logrus.New()

	db, err := localPool.NewLocalDocDB()
	if err != nil {
		log.Errorf("failed to connect to mongodb")
		panic(err)
	}

	fmt.Println("Successfully connected!", db.Name())

	// Repositories
	studyGroupRepo := mongodb.NewStudyGroupRepository(&mongodb.DocDB{
		Pool: db,
	})

	// Services
	studyGroupService := services.NewStudyGroupService(studyGroupRepo)

	// API init
	API := http.NewApi(studyGroupService)

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
