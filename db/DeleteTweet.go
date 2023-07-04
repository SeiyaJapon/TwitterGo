package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string) error {
	ctx := context.TODO()
	db := MongoCN.Database(Database)
	tweetCollection := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := tweetCollection.DeleteOne(ctx, condition)

	return err
}
