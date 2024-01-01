package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vitocom.com/community/db"
)

type Community struct {
	ID              primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	CommunityAdmins []primitive.ObjectID `json:"communityAdmins,omitempty" bson:"communityAdmins,omitempty"`
	Creator         primitive.ObjectID   `json:"creator,omitempty" bson:"creator,omitempty"`
	DateCreated     primitive.DateTime   `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
	Description     string               `json:"description,omitempty" bson:"description,omitempty"`
	Name            string               `json:"name,omitempty" bson:"name,omitempty"`
	OnApproved      bool                 `json:"onApproved,omitempty" bson:"onApproved,omitempty"`
	Type            string               `json:"type,omitempty" bson:"type,omitempty"`
	SubsribersCount int                  `json:"subsribersCount,omitempty" bson:"subsribersCount,omitempty"`
}

func (cm Community) Save() (*Community, error) {
	var communitiesCollection = db.MongoClient.Database("vitocom").Collection("communities")
	newDoc := bson.M{
		"name":            cm.Name,
		"description":     cm.Description,
		"onApproved":      false,
		"creator":         cm.Creator,
		"type":            cm.Type,
		"subsribersCount": 0,
		"dateCreated":     primitive.NewDateTimeFromTime(time.Now()),
	}

	result, err := communitiesCollection.InsertOne(context.TODO(), newDoc)

	if err != nil {
		return nil, err
	}
	cm.ID = result.InsertedID.(primitive.ObjectID)
	return &cm, nil
}

func GetAllCommunities() ([]Community, error) {
	var communitiesCollection = db.MongoClient.Database("vitocom").Collection("communities")

	// retrieve all the documents that match the filter
	cursor, err := communitiesCollection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		//panic(err)
		return nil, err
	}

	var communities []Community

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var cmt Community
		cursor.Decode(&cmt)
		communities = append(communities, cmt)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return communities, nil

}
