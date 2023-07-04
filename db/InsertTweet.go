package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.StoreTweet) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(Database)
	tweetsCollection := db.Collection("tweet")

	register := bson.M{
		"userid":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := tweetsCollection.InsertOne(ctx, register)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
