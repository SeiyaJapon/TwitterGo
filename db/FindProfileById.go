package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfileById(ID string) (models.User, error) {
	ctx := context.TODO()
	db := MongoCN.Database(Database)
	usersCollection := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	searchCondition := bson.M{
		"_id": objID,
	}

	err := usersCollection.FindOne(ctx, searchCondition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		return profile, err
	}

	return profile, nil
}
