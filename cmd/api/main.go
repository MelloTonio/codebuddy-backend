package main

import (
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	localPool "github.com/mellotonio/desafiogo"

	challengeService "github.com/mellotonio/desafiogo/app/domain/challenges/services"
	profileService "github.com/mellotonio/desafiogo/app/domain/profiles/services"
	groupService "github.com/mellotonio/desafiogo/app/domain/studyGroups/services"
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
	challengeRepo := mongodb.NewChallengeRepository(&mongodb.DocDB{
		Pool: db,
	})
	profileRepo := mongodb.NewProfileRepository(&mongodb.DocDB{
		Pool: db,
	})

	// Services
	profileService := profileService.NewProfileService(profileRepo)
	studyGroupService := groupService.NewStudyGroupService(studyGroupRepo, profileService)
	challengeService := challengeService.NewChallengeService(challengeRepo)

	// API init
	API := http.NewApi(studyGroupService, challengeService, profileService)

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
