package mongodb

import (
	"context"
	"log"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const challengeCollection = "challengeCollection"

type ChallengeRepo struct {
	DocDb *DocDB
}

func NewChallengeRepository(DocDb *DocDB) challenges.Repository {
	return &ChallengeRepo{
		DocDb: DocDb,
	}
}

func (cr *ChallengeRepo) CreateChallenge(ctx context.Context, challenge challenges.Challenge) error {
	challengeToBeInserted := bson.M{
		"name":       challenge.Name,
		"text":       challenge.Text,
		"group":      challenge.Group,
		"input":      challenge.Input,
		"output":     challenge.Output,
		"answers":    challenge.Answer,
		"difficulty": challenge.Difficult,
	}

	logrus.Infof("inserting %+v", challengeToBeInserted)

	_, err := cr.DocDb.Pool.Collection(challengeCollection).InsertOne(ctx, challengeToBeInserted)

	logrus.Infof("success %+v", challengeToBeInserted)
	return err
}

func (cr *ChallengeRepo) UpdateChallenge(ctx context.Context, challenge challenges.Challenge) error {
	filter := bson.M{
		"group": challenge.Group,
		"name":  challenge.Name,
	}

	newAnswer := challenges.Answer{
		AlumniName: challenge.Answer[0].AlumniName,
		Text:       challenge.Answer[0].Text,
	}

	update := bson.M{
		"$push": bson.M{
			"answer": newAnswer,
		},
	}

	_, err := cr.DocDb.Pool.Collection(challengeCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (cr *ChallengeRepo) ListChallengesByGroup(ctx context.Context, groupName string) ([]challenges.Challenge, error) {
	var dbChallenges []challenges.Challenge
	cursor, err := cr.DocDb.Pool.Collection(challengeCollection).Find(ctx, bson.M{
		"group": groupName,
	})
	if err != nil {
		logrus.Errorf("failed to get study groups: %s", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var challenge challenges.Challenge
		err := cursor.Decode(&challenge)
		if err != nil {
			logrus.Errorf("failed to decode study group: %s", err.Error())
			return nil, err
		}
		dbChallenges = append(dbChallenges, challenge)
	}
	if err := cursor.Err(); err != nil {
		logrus.Errorf("cursor error: %s", err.Error())
		return nil, err
	}

	logrus.Infof("got %+v", dbChallenges)

	return dbChallenges, nil
}
