package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRegister(user models.User, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(Database)
	usersCollection := db.Collection("users")
	register := make(map[string]interface{})

	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	if len(user.Surname) > 0 {
		register["surname"] = user.Surname
	}

	register["birthday"] = user.Birthday

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	if len(user.Bio) > 0 {
		register["bio"] = user.Bio
	}

	if len(user.Location) > 0 {
		register["location"] = user.Location
	}

	if len(user.Website) > 0 {
		register["website"] = user.Website
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := usersCollection.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
