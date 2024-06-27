package mongodb

import (
	"context"
	"errors"
	"log"

	"github.com/mellotonio/desafiogo/app/domain/profiles"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const profileCollection = "profileCollection"

type ProfileRepo struct {
	DocDb *DocDB
}

func NewProfileRepository(DocDb *DocDB) profiles.Repository {
	return &ProfileRepo{
		DocDb: DocDb,
	}
}

func (cr *ProfileRepo) CreateProfile(ctx context.Context, profile profiles.Profile) error {
	profileToBeInserted := bson.M{
		"username":     profile.Username,
		"password":     profile.Password,
		"groups":       []string{},
		"profile_type": profile.ProfileType,
	}

	logrus.Infof("inserting %+v", profileToBeInserted)

	_, err := cr.DocDb.Pool.Collection(profileCollection).InsertOne(ctx, profileToBeInserted)

	logrus.Infof("success %+v", profileToBeInserted)
	return err
}

func (cr *ProfileRepo) GetAllProfiles(ctx context.Context) ([]profiles.Profile, error) {
	var results []profiles.Profile
	cur, err := cr.DocDb.Pool.Collection(profileCollection).Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem profiles.Profile
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, profiles.Profile{
			Username: elem.Username,
		})
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	logrus.Infof("success %+v", results)
	return results, err
}

func (cr *ProfileRepo) UpdateProfileByUsername(ctx context.Context, username string, groupName string) error {
	filter := bson.M{
		"username": username,
	}

	logrus.Infof("updating %s to %s", username, groupName)
	update := bson.M{
		"$push": bson.M{
			"groups": groupName,
		},
	}

	_, err := cr.DocDb.Pool.Collection(profileCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (cr *ProfileRepo) GetProfile(ctx context.Context, profile profiles.Profile) (profiles.Profile, error) {
	var storedProfile profiles.Profile
	filter := bson.M{"username": profile.Username}

	err := cr.DocDb.Pool.Collection(profileCollection).FindOne(ctx, filter).Decode(&storedProfile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Errorf("profile not found: %v", err)
			return profiles.Profile{}, errors.New("profile not found")
		}
		logrus.Errorf("error finding profile: %v", err)
		return profiles.Profile{}, err
	}

	return storedProfile, nil
}
