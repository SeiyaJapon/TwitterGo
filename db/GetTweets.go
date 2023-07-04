package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(Database)
	tweetsCollection := db.Collection("tweet")

	var result []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()

	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := tweetsCollection.Find(ctx, condition, options)

	if err != nil {
		return result, false
	}

	for cursor.Next(ctx) {
		var register models.ReturnTweets
		err := cursor.Decode(&register)

		if err != nil {
			return result, false
		}

		result = append(result, &register)
	}

	return result, true
}
