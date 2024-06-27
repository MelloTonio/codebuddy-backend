package mongodb

import (
	"context"
	"log"

	studygroups "github.com/mellotonio/desafiogo/app/domain/studyGroups"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const studyGroupsCollection = "studyGroupsCollection"

type DocDB struct {
	Pool *mongo.Database
}

type StudyGroupRepo struct {
	DocDb *DocDB
}

func NewStudyGroupRepository(DocDb *DocDB) studygroups.Repository {
	return &StudyGroupRepo{
		DocDb: DocDb,
	}
}

func (cr *StudyGroupRepo) SaveStudyGroup(ctx context.Context, stdGroup studygroups.StudyGroup) error {

	existingStudyGroup := cr.DocDb.Pool.Collection(studyGroupsCollection).FindOne(ctx, bson.M{"name": stdGroup.Name})
	if existingStudyGroup.Err() == nil {
		logrus.Infof("group already exist")
		return nil
	} else if existingStudyGroup.Err() != mongo.ErrNoDocuments {
		return existingStudyGroup.Err()
	}

	studyGroupToBeInserted := bson.M{
		"name":        stdGroup.Name,
		"students":    stdGroup.Students,
		"subject":     stdGroup.Subject,
		"description": stdGroup.Description,
		"warnings":    []string{},
	}

	logrus.Infof("inserting %+v", studyGroupToBeInserted)

	_, err := cr.DocDb.Pool.Collection(studyGroupsCollection).InsertOne(ctx, studyGroupToBeInserted)

	logrus.Infof("success %+v", studyGroupToBeInserted)
	return err
}

func (rr *StudyGroupRepo) GetStudyGroupsByStudent(ctx context.Context, studentName string) ([]studygroups.StudyGroup, error) {
	var studyGroups []studygroups.StudyGroup

	// Query to find study groups where the student name is in the "students" array
	cursor, err := rr.DocDb.Pool.Collection(studyGroupsCollection).Find(ctx, bson.M{
		"students": studentName,
	})
	if err != nil {
		logrus.Errorf("failed to get study groups: %s", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var studyGroup studygroups.StudyGroup
		err := cursor.Decode(&studyGroup)
		if err != nil {
			logrus.Errorf("failed to decode study group: %s", err.Error())
			return nil, err
		}
		studyGroups = append(studyGroups, studyGroup)
	}

	if err := cursor.Err(); err != nil {
		logrus.Errorf("cursor error: %s", err.Error())
		return nil, err
	}

	logrus.Infof("got %+v", studyGroups)

	return studyGroups, nil
}

func (rr *StudyGroupRepo) GetStudyGroupDetails(ctx context.Context, studyGroupName string) (studygroups.StudyGroup, error) {
	var studyGroup studygroups.StudyGroup

	err := rr.DocDb.Pool.Collection(studyGroupsCollection).FindOne(ctx, bson.M{"name": studyGroupName}).Decode(&studyGroup)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No study group found with the name: %s", studyGroupName)
			return studygroups.StudyGroup{}, nil // or return an appropriate error
		}
		log.Printf("Failed to get study group details: %v", err)
		return studygroups.StudyGroup{}, err
	}

	return studyGroup, nil
}

func (rr *StudyGroupRepo) SaveWarning(ctx context.Context, studyGroupName, warnMessage string) error {
	filter := bson.M{"name": studyGroupName}

	update := bson.M{
		"$push": bson.M{
			"warnings": warnMessage,
		},
	}

	logrus.Infof("%+v", update)

	result, err := rr.DocDb.Pool.Collection(studyGroupsCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Failed to save warning message: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		log.Printf("No study group found with the name: %s", studyGroupName)
		return mongo.ErrNoDocuments
	}

	log.Printf("Warning message added to study group: %s", studyGroupName)
	return nil
}

func (rr *StudyGroupRepo) ListGroupStudents(ctx context.Context, groupName string) ([]studygroups.StudyGroup, error) {
	var studyGroups []studygroups.StudyGroup

	// Query to find study groups where the student name is in the "students" array
	cursor, err := rr.DocDb.Pool.Collection(studyGroupsCollection).Find(ctx, bson.M{
		"name": groupName,
	})
	if err != nil {
		logrus.Errorf("failed to get study groups: %s", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var studyGroup studygroups.StudyGroup
		err := cursor.Decode(&studyGroup)
		if err != nil {
			logrus.Errorf("failed to decode study group: %s", err.Error())
			return nil, err
		}
		studyGroups = append(studyGroups, studyGroup)
	}

	if err := cursor.Err(); err != nil {
		logrus.Errorf("cursor error: %s", err.Error())
		return nil, err
	}

	logrus.Infof("got %+v", studyGroups)

	return studyGroups, nil
}

func (rr *StudyGroupRepo) GetWarnings(ctx context.Context, studyGroupName string) ([]string, error) {
	var studyGroup studygroups.StudyGroup

	filter := bson.M{"name": studyGroupName}

	err := rr.DocDb.Pool.Collection(studyGroupsCollection).FindOne(ctx, filter).Decode(&studyGroup)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No study group found with the name: %s", studyGroupName)
			return nil, nil
		}
		log.Printf("Failed to find study group: %v", err)
		return nil, err
	}

	logrus.Infof("%+v", studyGroup.Warnings)
	return studyGroup.Warnings, nil
}

func (rr *StudyGroupRepo) AddStudentsToGroup(ctx context.Context, studyGroupName string, students []string) error {
	filter := bson.M{"name": studyGroupName}

	update := bson.M{
		"$addToSet": bson.M{
			"students": bson.M{
				"$each": students,
			},
		},
	}

	logrus.Infof("updating study group %s with students %+v", studyGroupName, students)

	result, err := rr.DocDb.Pool.Collection(studyGroupsCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Failed to add students to study group: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		log.Printf("No study group found with the name: %s", studyGroupName)
		return mongo.ErrNoDocuments
	}

	log.Printf("Students added to study group: %s", studyGroupName)
	return nil
}
